package impl

import (
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/xbklyn/getgoal-app/common"
	"github.com/xbklyn/getgoal-app/config"
	"github.com/xbklyn/getgoal-app/entity"
	"github.com/xbklyn/getgoal-app/model"
	repository "github.com/xbklyn/getgoal-app/repository"
	"github.com/xbklyn/getgoal-app/service"
	"golang.org/x/crypto/argon2"
)

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

var p = &params{
	memory:      64 * 1024, // 64 MB
	iterations:  3,
	parallelism: 1,
	saltLength:  16,
	keyLength:   32,
}

func NewAuthServiceImpl(userRepo repository.UserRepo, mailer service.MailerService) service.AuthService {
	return &AuthServiceImpl{userRepo, mailer}
}

type AuthServiceImpl struct {
	UserRepo repository.UserRepo
	Mailer   service.MailerService
}

// Verify implements service.AuthService.
func (service *AuthServiceImpl) Verify(request model.VerifyRequest) error {
	//find user by email
	user, err := service.UserRepo.FindUserByEmail(request.Email)
	if err != nil {
		return err
	}
	//check if user is already verified
	if user.EmailValidationStatusID == 1 {
		return errors.New("user is already verified")
	}
	//check if verification code is valid
	if user.ConfirmationToken != request.Code {
		return errors.New("invalid verification code")
	}
	//check if verification code is expired
	if time.Since(user.TokenGenerationTime).Hours() > 24 {
		return errors.New("verification code is expired")
	}
	//update user email_validation = 1
	user.EmailValidationStatusID = 1
	err = service.UserRepo.Update(user.UserID, user)
	if err != nil {
		return err
	}
	return nil
}

// SignUp implements service.AuthService.
func (service *AuthServiceImpl) SignUp(request model.SignUpRequest) (useEntityr entity.UserAccount, err error) {
	validateErr := common.Validate(request)
	if validateErr != nil {
		return entity.UserAccount{}, validateErr
	}
	//check existing user with email
	user, _ := service.UserRepo.FindUserByEmail(request.Email)
	if user.UserID != 0 {
		return entity.UserAccount{}, errors.New("this email is already registered")
	}
	//gen passwod
	hashed, encodedHash, err := generateHashFromPassword(request.Password, p)
	if err != nil {
		return entity.UserAccount{}, err
	}

	//gen verification code
	verificationCode := generateVerificationCode(6)
	//save user with email_validation = 2
	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Email = request.Email

	user.PasswordHash = hashed
	user.PasswordSalt = encodedHash

	user.ConfirmationToken = verificationCode
	user.TokenGenerationTime = time.Now()
	user.EmailValidationStatusID = 2

	err = service.UserRepo.Save(&user)
	if err != nil {
		return entity.UserAccount{}, err
	}

	// send email
	data := model.EmailTemplateData{
		VerificationCode: verificationCode,
	}
	if err := service.Mailer.SendEmail([]string{user.Email}, config.VERIFICATION_SUBJECT+verificationCode, config.VERIFICATION_TEMPLATE, data); err != nil {
		return entity.UserAccount{}, err
	}
	return user, nil
}

func generateHashFromPassword(password string, p *params) (hashed string, encodedHash string, err error) {
	salt, err := generateRandomBytes(p.saltLength)
	if err != nil {
		return "", "", err
	}

	hash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	fmt.Print(string(hash))
	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Return a string using the standard encoded hash representation.
	encodedHash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, p.memory, p.iterations, p.parallelism, b64Salt, b64Hash)

	return b64Hash, encodedHash, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func verifyPassword(password, encodedHash string) (match bool, err error) {
	// Extract the parameters, salt and derived key from the encoded password
	// hash.
	p, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		return false, err
	}

	// Derive the key from the other password using the same parameters.
	otherHash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	// Check that the contents of the hashed passwords are identical. Note
	// that we are using the subtle.ConstantTimeCompare() function for this
	// to help prevent timing attacks.
	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}
	return false, nil
}

func generateVerificationCode(length int) string {
	// Define the character set for the verification code (numbers only)
	charSet := "0123456789"

	// Initialize the verification code string
	verificationCode := make([]byte, length)

	// Populate the verification code with random characters from the character set
	for i := range verificationCode {
		verificationCode[i] = charSet[rand.Intn(len(charSet))]
	}

	return string(verificationCode)
}
func decodeHash(encodedHash string) (p *params, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, errors.New("the encoded hash is not in the correct format")
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, errors.New("incompatible version of argon2")
	}

	p = &params{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.saltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.keyLength = uint32(len(hash))

	return p, salt, hash, nil
}

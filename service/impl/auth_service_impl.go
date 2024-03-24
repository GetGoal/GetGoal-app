package impl

import (
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
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
	return &AuthServiceImpl{userRepo, mailer, make(map[string]bool), common.Claims{}}
}

type AuthServiceImpl struct {
	UserRepo         repository.UserRepo
	Mailer           service.MailerService
	BlackListedToken map[string]bool
	Claims           common.Claims
}

// ResetPassword implements service.AuthService.
func (service *AuthServiceImpl) ResetPassword(request model.ResetPasswordRequest) error {
	//find user by email
	user, _ := service.UserRepo.FindUserByEmail(request.Email)
	if user.UserID == 0 {
		return errors.New("email not found")
	}
	//check if len of provider is 0
	if len(user.ExternalProvider) != 0 {
		return errors.New("user is registered with external provider, please login using external provider instead")
	}
	//gen new access token
	access, _, err := common.GenerateToken(user)
	if err != nil {
		return err
	}

	resetLink := fmt.Sprintf("%s?token=%s", config.GetConfig().Mailer.BaseURL, access)
	//send email
	data := model.ResetPasswordTemplateData{
		ResetLink: resetLink,
	}

	if err := service.Mailer.SendEmail([]string{user.Email}, config.RESET_PASSWORD_SUBJECT, config.RESET_PASSWORD_TEMPLATE, data); err != nil {
		return err
	}

	return nil
}

// IsTokenBlacklisted implements service.AuthService.
func (service *AuthServiceImpl) IsTokenBlacklisted(tokenString string) bool {
	// Check if token is blacklisted
	_, blacklisted := service.BlackListedToken[tokenString]
	return blacklisted
}

// SignOut implements service.AuthService.
func (service *AuthServiceImpl) SignOut(token string) error {
	service.BlackListedToken[token] = true
	return nil
}

// SignInWithGoogle implements service.AuthService.
func (service *AuthServiceImpl) ExternalSignIn(request model.ProviderSignInRequest) (accessToken string, refreshToken string, err error) {
	// validate request
	if err := common.Validate(request); err != nil {
		return "", "", err
	}

	switch request.Provider {
	case "google":
		accessToken, refreshToken, err = service.signInWithGoogle(request.Google)
		if err != nil {
			return "", "", err
		}
		return accessToken, refreshToken, nil
	default:
		return "", "", errors.New("invalid provider")
	}
}

// SignIn implements service.AuthService.
func (service *AuthServiceImpl) SignIn(request model.Credentials) (string, string, error) {
	if err := common.Validate(request); err != nil {
		return "", "", err
	}
	//find user by email
	user, _ := service.UserRepo.FindUserByEmail(request.Email)
	if user.UserID == 0 {
		return "", "", errors.New("user not found")
	}
	//check if user is verified
	if user.EmailValidationStatusID != 1 {
		return "", "", errors.New("user is not verified")
	}
	//check if password is correct
	match, err := verifyPassword(request.Password, user.PasswordSalt)
	if err != nil {
		return "", "", err
	}
	if !match {
		return "", "", errors.New("invalid password")
	}
	//generate access token
	access, refresh, err := common.GenerateToken(user)
	if err != nil {
		return "", "", err
	}

	return access, refresh, nil
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

func (service *AuthServiceImpl) signInWithGoogle(request model.GoogleSignInRequest) (accessToken string, refreshToken string, err error) {
	// validate request
	if err := common.Validate(request); err != nil {
		return "", "", err
	}
	// find user by email
	user, _ := service.UserRepo.FindUserByEmail(request.Email)
	if user.UserID == 0 {
		log.Default().Println("Create new user with google credential")
		// create user
		user.FirstName = request.DisplayName
		user.Email = request.Email
		user.EmailValidationStatusID = 1

		//display picture
		// user. = request.PhotoURL
		user.ExternalProvider = append(user.ExternalProvider, entity.ExternalProvider{
			ExternalProviderID: 1,
			ProviderName:       "google",
		})
		err := service.UserRepo.Save(&user)
		if err != nil {
			return "", "", err
		}
		// generate access token
		access, refresh, err := common.GenerateToken(user)
		if err != nil {
			return "", "", err
		}
		return access, refresh, nil
	}

	// check if provider is google
	var hasGoogle bool
	for _, provider := range user.ExternalProvider {
		if provider.ProviderName == "google" {
			hasGoogle = true
		}
	}
	if !hasGoogle {
		return "", "", errors.New("user already registered with getgoal account, try logging in with this email")
	}
	// check if user is verified
	if user.EmailValidationStatusID != 1 {
		return "", "", errors.New("user is not verified")
	}
	// generate access token
	access, refresh, err := common.GenerateToken(user)
	if err != nil {
		return "", "", err
	}
	return access, refresh, nil
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

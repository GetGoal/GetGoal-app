package impl

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/xbklyn/getgoal-app/common"
	"github.com/xbklyn/getgoal-app/config"
	"github.com/xbklyn/getgoal-app/entity"
	"github.com/xbklyn/getgoal-app/model"
	repository "github.com/xbklyn/getgoal-app/repository"
	"github.com/xbklyn/getgoal-app/service"
	"github.com/zhenghaoz/gorse/client"
)

func NewAuthServiceImpl(userRepo repository.UserRepo, programRepo repository.ProgramRepo, mailer service.MailerService, gorse client.GorseClient) service.AuthService {
	return &AuthServiceImpl{userRepo, programRepo, mailer, make(map[string]bool), common.Claims{}, gorse}
}

type AuthServiceImpl struct {
	UserRepo         repository.UserRepo
	ProgramRepo      repository.ProgramRepo
	Mailer           service.MailerService
	BlackListedToken map[string]bool
	Claims           common.Claims
	Gorse            client.GorseClient
}

// VerifyPasswordReset implements service.AuthService.
func (service *AuthServiceImpl) VerifyPasswordReset(request model.VerifyResetRequest) error {
	//find user by email
	user, _ := service.UserRepo.FindUserByEmail(request.Email)
	if user.UserID == 0 {
		return errors.New("email not found")
	}
	//check if len of provider is 0
	if len(user.ExternalProvider) != 0 {
		return errors.New("user is registered with external provider, please login using external provider instead")
	}
	//check if verification code is valid
	if user.PasswordRecoveryToken != request.Code {
		return errors.New("invalid verification code")
	}
	//check if verification code is expired
	if time.Since(user.TokenGenerationTime).Hours() > 24 {
		return errors.New("verification code is expired")
	}
	return nil
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

	verificationCode := generateVerificationCode(6)
	//update user with new verification code
	user.PasswordRecoveryToken = verificationCode
	user.TokenGenerationTime = time.Now()
	err := service.UserRepo.Update(user.UserID, user)
	if err != nil {
		return err
	}
	//send email
	data := model.ResetPasswordTemplateData{
		ResetCode: verificationCode,
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
	match, err := common.VerifyPassword(request.Password, user.PasswordSalt)
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
func (service *AuthServiceImpl) Verify(request model.VerifyRequest) (accessToken string, refreshToken string, err error) {
	//find user by email
	user, err := service.UserRepo.FindUserByEmail(request.Email)
	if err != nil {
		return "", "", err
	}
	//check if user is already verified
	if user.EmailValidationStatusID == 1 {
		return "", "", errors.New("user is already verified")
	}
	//check if verification code is valid
	if user.ConfirmationToken != request.Code {
		return "", "", errors.New("invalid verification code")
	}
	//check if verification code is expired
	if time.Since(user.TokenGenerationTime).Hours() > 24 {
		return "", "", errors.New("verification code is expired")
	}
	//update user email_validation = 1
	user.EmailValidationStatusID = 1
	err = service.UserRepo.Update(user.UserID, user)
	if err != nil {
		return "", "", err
	}

	accessToken, refreshToken, err = common.GenerateToken(user)
	if err != nil {
		return "", "", err

	}
	return accessToken, refreshToken, nil
}

// SignUp implements service.AuthService.
func (service *AuthServiceImpl) SignUp(request model.SignUpRequest) (useEntityr entity.UserAccount, err error) {
	validateErr := common.Validate(request)
	if validateErr != nil {
		return entity.UserAccount{}, validateErr
	}
	//check existing user with email
	existedUser, _ := service.UserRepo.FindUserByEmail(request.Email)
	if existedUser.UserID != 0 {
		return entity.UserAccount{}, errors.New("this email is already registered")
	}
	//gen passwod
	hashed, encodedHash, err := common.GenerateHashFromPassword(request.Password)
	if err != nil {
		return entity.UserAccount{}, err
	}
	// log.Default().Printf("json data: %s", jsonData)
	//gen verification code
	verificationCode := generateVerificationCode(6)
	//save user with email_validation = 2

	text, _ := json.Marshal(request.Labels)
	newUser := entity.UserAccount{
		FirstName:               request.FirstName,
		LastName:                request.LastName,
		Email:                   request.Email,
		Labels:                  string(text),
		PasswordHash:            hashed,
		PasswordSalt:            encodedHash,
		ConfirmationToken:       verificationCode,
		TokenGenerationTime:     time.Now(),
		EmailValidationStatusID: 2,
	}

	err = service.UserRepo.Save(&newUser)
	if err != nil {
		return entity.UserAccount{}, err
	}

	// send email
	data := model.EmailTemplateData{
		VerificationCode: verificationCode,
	}
	if err := service.Mailer.SendEmail([]string{newUser.Email}, config.VERIFICATION_SUBJECT+verificationCode, config.VERIFICATION_TEMPLATE, data); err != nil {
		return entity.UserAccount{}, err
	}
	_, gErr := service.Gorse.InsertUser(context.TODO(), client.User{
		UserId: strconv.Itoa(int(newUser.UserID)),
		Labels: request.Labels,
	})
	if gErr != nil {
		return entity.UserAccount{}, gErr
	}

	programs, err := service.ProgramRepo.FindProgramByLabelWithLimits(request.Labels, 10)
	if err != nil {
		return entity.UserAccount{}, err
	}

	var feedbacks []client.Feedback
	for _, program := range programs {
		feedbacks = append(feedbacks, client.Feedback{
			FeedbackType: "view_program",
			UserId:       strconv.Itoa(int(newUser.UserID)),
			ItemId:       strconv.Itoa(int(program.ProgramID)),
			Timestamp:    time.Now().Format("2006-01-02"),
		})
	}
	log.Default().Printf("feedbacks numbers: %d", len(feedbacks))
	_, gErr = service.Gorse.InsertFeedback(context.TODO(), feedbacks)
	if gErr != nil {
		return entity.UserAccount{}, gErr
	}
	return newUser, nil
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
		user.Labels = "[]"

		//display picture
		user.ImageURL = request.PhotoURL
		user.ExternalProvider = append(user.ExternalProvider, entity.ExternalProvider{
			ExternalProviderID: 1,
			ProviderName:       "google",
		})
		err := service.UserRepo.Save(&user)
		if err != nil {
			return "", "", err
		}

		_, gErr := service.Gorse.InsertUser(context.TODO(), client.User{
			UserId: strconv.Itoa(int(user.UserID)),
			Labels: []string{},
		})
		if gErr != nil {
			return "", "", gErr
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

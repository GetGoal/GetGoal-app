package common

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"github.com/xbklyn/getgoal-app/config"
	"github.com/xbklyn/getgoal-app/entity"
)

type Claims struct {
	Email     string `json:"email"`
	UserID    uint64 `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	jwt.StandardClaims
}

type RefreshClaims struct {
	Email     string `json:"email"`
	UserID    uint64 `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	jwt.StandardClaims
}

func GenerateToken(user entity.UserAccount) (string, string, error) {
	accessToken, err := generateAccessToken(user)
	if err != nil {
		return "", "", err
	}

	refreshTokenString, err := generateRefreshToken(user)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshTokenString, nil
}

func generateAccessToken(user entity.UserAccount) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		UserID:    user.UserID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(config.GetConfig().JwtKeys.AccessSecret))
	if err != nil {
		return "", err
	}
	return accessToken, nil
}
func generateAccessTokenForResetPassword(user entity.UserAccount) (string, error) {
	expirationTime := time.Now().Add(10 * time.Minute)
	claims := &Claims{
		UserID:    user.UserID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(config.GetConfig().JwtKeys.AccessSecret))
	if err != nil {
		return "", err
	}
	return accessToken, nil

}

func generateRefreshToken(user entity.UserAccount) (string, error) {
	refreshExpirationTime := time.Now().Add(7 * 24 * time.Hour)
	refreshClaims := &RefreshClaims{
		UserID:    user.UserID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshExpirationTime.Unix(),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(config.GetConfig().JwtKeys.RefreshSecret))
	if err != nil {
		return "", err
	}
	return refreshTokenString, nil
}
func RefreshTokens(refreshTokenString string) (string, string, error) {
	// Validate Refresh Token
	claims, err := validateRefreshToken(refreshTokenString)
	if err != nil {
		return "", "", err
	}

	// Check if refresh token is expired
	if time.Now().Unix() > claims.ExpiresAt {
		return "", "", errors.New("refresh token is expired")
	}

	// Generate new Access Token
	accessToken, err := generateAccessToken(claims.User)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshTokenString, nil
}

func validateRefreshToken(tokenString string) (*RefreshClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().JwtKeys.RefreshSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*RefreshClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}

func ValidateAccessToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().JwtKeys.AccessSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}

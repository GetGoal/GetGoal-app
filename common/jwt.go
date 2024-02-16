package common

import (
	"time"

	"github.com/golang-jwt/jwt"
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
	jwt.StandardClaims
}

func GenerateToken(user entity.UserAccount) (string, string, error) {
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
		return "", "", err
	}

	refreshExpirationTime := time.Now().Add(7 * 24 * time.Hour)
	refreshClaims := &RefreshClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshExpirationTime.Unix(),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(config.GetConfig().JwtKeys.RefreshSecret))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshTokenString, nil
}

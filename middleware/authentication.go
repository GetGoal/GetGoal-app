package middleware

// auth/middleware/auth_middleware.go

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/xbklyn/getgoal-app/common"
	"github.com/xbklyn/getgoal-app/model"
	"github.com/xbklyn/getgoal-app/service/impl"
)

func JWTAuthMiddleware(service *impl.AuthServiceImpl, jwtKey []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract token from request headers
		tokenString := extractToken(c)
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, model.GeneralResponse{
				Code:    http.StatusUnauthorized,
				Message: "No token provided",
				Data:    nil,
				Error:   "No token provided",
			})
			c.Abort()
			return
		}

		// Check if token is blacklisted
		if service.IsTokenBlacklisted(tokenString) {
			c.JSON(http.StatusUnauthorized, model.GeneralResponse{
				Code:    http.StatusUnauthorized,
				Message: "Invalid token",
				Data:    nil,
				Error:   "Token is already blacklisted, please sign in again",
			})
			c.Abort()
			return
		}

		// Validate token
		claims, err := validateAccessToken(tokenString, jwtKey, service)
		if err != nil {
			log.Default().Println("Error validating token: ", err)
			if strings.Contains(err.Error(), "token is expired by") {
				// Token is expired, try refreshing
				log.Default().Println("Token is expired, try refreshing")
				refreshToken := c.Request.Header.Get("RefreshToken")
				if refreshToken == "" {
					c.AbortWithStatusJSON(http.StatusUnauthorized, model.GeneralResponse{
						Code:    http.StatusUnauthorized,
						Message: "Refresh token required",
						Data:    nil,
						Error:   "Token is expired, Need refresh token to proceed",
					})
					return
				}

				newAccessToken, newRefreshToken, err := common.RefreshTokens(refreshToken)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusUnauthorized, model.GeneralResponse{
						Code:    http.StatusUnauthorized,
						Message: "Cannot refresh token",
						Data:    nil,
						Error:   "Invalid refresh token",
					})
					return
				}
				newClaims, err := validateAccessToken(newAccessToken, jwtKey, service)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusUnauthorized, model.GeneralResponse{
						Code:    http.StatusUnauthorized,
						Message: "Invalid token",
						Data:    nil,
						Error:   err.Error(),
					})
					return
				}
				c.Set("claims", newClaims)
				c.Set("access_token", newAccessToken)
				// Set new access token in response headers
				c.Writer.Header().Set("Authorization", "Bearer "+newAccessToken)
				// Optionally, set new refresh token in response headers
				c.Writer.Header().Set("RefreshToken", newRefreshToken)
				log.Default().Printf("New access token: %s", newAccessToken)
				log.Default().Printf("New refresh token: %s", newRefreshToken)
				log.Default().Printf("New claims: %v", newClaims)
				// Proceed with the request using the new access token
				c.Next()
				return
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.GeneralResponse{
				Code:    http.StatusUnauthorized,
				Message: "Invalid token",
				Data:    nil,
				Error:   err.Error(),
			})
			return
		}
		c.Set("claims", claims)
		c.Set("access_token", tokenString)

		// Call the next handler
		c.Next()
	}
}

func extractToken(c *gin.Context) string {
	// Extract token from Authorization header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return ""
	}
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}
	return parts[1]
}

func validateAccessToken(tokenString string, jwtKey []byte, service *impl.AuthServiceImpl) (*common.Claims, error) {
	// Validate token
	token, err := jwt.ParseWithClaims(tokenString, &service.Claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {

		return nil, err
	}

	// Check if token is expired
	claims, ok := token.Claims.(*common.Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

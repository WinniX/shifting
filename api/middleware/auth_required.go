package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/winnix/shifting/api/config"
	"github.com/winnix/shifting/api/service"
)

const BEARER_PREFIX = "Bearer "
const AUTH_HEADER = "Authorization"

func AuthRequired(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(AUTH_HEADER)
		if len(authHeader) == 0 {
			fmt.Print("AuthRequired: authorization header is missing\n")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		authToken := strings.Replace(authHeader, BEARER_PREFIX, "", 1)
		token, err := service.NewJwtService(cfg.JwtSigningKey).ValidateToken(authToken)
		if err != nil {
			fmt.Printf("AuthRequired: invalid token: %s\n", err.Error())
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			fmt.Print("AuthRequired: invalid token\n")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		c.Set("accountCode", claims["account_code"])
		c.Set("propertyId", claims["property_id"])
		c.Set("subjectId", claims["sub"])
		c.Next()
	}
}

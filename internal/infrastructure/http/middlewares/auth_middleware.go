package middlewares

import (
	"fmt"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/usecases"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

//AuthorizeJWTMiddleware Jwt middleware
func AuthorizeJWTMiddleware(a usecases.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := a.ValidateToken(c.Request)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}

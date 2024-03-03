package middleware

import (
	"akozadaev/go_reports/db/authentication"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthenticationMiddleware(auth authentication.Authentication) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := auth.Authenticate(c.Request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}

		c.Set(authentication.User, user)
		c.Next()
	}
}

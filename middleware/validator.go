package middleware

import (
	"gintest/auth"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  err.Error(),
			})
			c.Abort()
			return
		}

		token, err := auth.ExtractTokenData(c.Request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Result": "NG",
				"message":  err,
			})
			return
		}

		log.Printf("permission : %v", token.Permission)
		for _, permission := range token.Permission.ListPermission {
			if permission.PermissionName == strings.Replace(c.Request.RequestURI, "/", "", 1) {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"Result": "NG",
			"message":  "You doesn't have access to this API",
		})
		c.Abort()
	}
}
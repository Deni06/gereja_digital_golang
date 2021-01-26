package controller

import (
	"gintest/auth"
	"gintest/dto"
	"gintest/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	user, err := repository.Login(dto.Login{Username: c.PostForm("username"), Password:c.PostForm("password")})
	if err != nil {
		log.Printf("error : %v", err)
		result = gin.H{
			"Result": "NG",
			"message":  err.Error(),
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	permission, err := repository.GetPermission(dto.Permission{UserID: user.UserID})
	if err != nil {
		log.Printf("error : %v", err)
		result = gin.H{
			"Result": "NG",
			"message":  err.Error(),
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}
		claims := jwt.MapClaims{}
		claims["authorized"] = true
		claims["user_id"] = user.UserID
		claims["username"] = user.Username
		claims["app_id"] = user.AppID
		claims["app_name"] = user.AppName
		claims["subscription_type"] = user.SubscriptionType
		claims["status_design"] = user.StatusDesign
		claims["permission"] = permission
		claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
		token, err := auth.CreateToken(claims)
		if err != nil {
			result = gin.H{
				"Result": "OK",
				"message":  "failed to generate token",
			}
			c.JSON(http.StatusInternalServerError, result)
			return
		}
		result = gin.H{
			"Result": "OK",
			"token": token,
		}

	c.JSON(http.StatusOK, result)
}
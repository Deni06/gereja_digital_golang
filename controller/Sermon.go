package controller

import (
	"gintest/auth"
	"gintest/dto"
	"gintest/repository"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetSermon(c *gin.Context) {
	token, err := auth.ExtractTokenData(c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if token.AppID == ""{
		c.JSON(http.StatusInternalServerError, "UnAuthorized")
		return
	}
		sermons, err := repository.GetAllSermon(dto.Sermon{AppID: token.AppID})
		if err != nil {
			log.Printf("error : %v", err)
			result = gin.H{
				"Result": "NG",
				"message":  "Failed Get Sermons",
			}

			c.JSON(http.StatusInternalServerError, result)
			return
		}

			result = gin.H{
				"Result": "OK",
				"message":  sermons,
			}

	c.JSON(http.StatusOK, result)
}
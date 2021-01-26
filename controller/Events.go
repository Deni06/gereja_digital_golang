package controller

import (
	"gintest/auth"
	"gintest/dto"
	"gintest/repository"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var result gin.H

func GetEvents(c *gin.Context) {
	token, err := auth.ExtractTokenData(c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Result": "NG",
			"message":  err,
		})
		return
	}
	if token.AppID == ""{
		c.JSON(http.StatusInternalServerError, gin.H{
			"Result": "NG",
			"message":  "UnAuthorized",
		})
		return
	}
		events, err := repository.GetAllEvents(dto.Events{AppID: token.AppID})
		if err != nil {
			log.Printf("error : %v", err)
			result = gin.H{
				"Result": "NG",
				"message":  "Failed Get Events",
			}
			c.JSON(http.StatusInternalServerError, result)
			return
		}
			result = gin.H{
				"Result": "OK",
				"message":  events,
			}

	c.JSON(http.StatusOK, result)
}
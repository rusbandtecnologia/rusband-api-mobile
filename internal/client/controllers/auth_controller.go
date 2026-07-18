package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rusbandtecnologia/rusband-api-mobile/internal/client/requests"
	"github.com/rusbandtecnologia/rusband-api-mobile/internal/client/services"
)

func Login(c *gin.Context) {

	var request requests.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Dados inválidos.",
		})

		return

	}

	client, err := services.Login(request.Email, request.Password)

	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "E-mail ou senha inválidos.",
		})

		return

	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"client":  client,
	})

}

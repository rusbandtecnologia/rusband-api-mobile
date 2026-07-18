package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rusbandtecnologia/rusband-api-mobile/internal/client/controllers"
	"github.com/rusbandtecnologia/rusband-api-mobile/internal/responses"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		responses.Success(
			c,
			"Rusband Mobile API",
			nil,
		)
	})

	client := router.Group("/client")
	{
		client.POST("/login", controllers.Login)
	}
}

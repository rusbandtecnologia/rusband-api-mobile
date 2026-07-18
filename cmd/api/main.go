package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/rusbandtecnologia/rusband-api-mobile/internal/config"
	"github.com/rusbandtecnologia/rusband-api-mobile/internal/database"
	"github.com/rusbandtecnologia/rusband-api-mobile/internal/routes"
)

func main() {

	config.LoadEnv()

	app := config.LoadApp()

	database.Connect(app)

	router := gin.Default()

	routes.RegisterRoutes(router)

	log.Printf("%s iniciada.", app.Name)

	router.Run(":" + app.Port)

}

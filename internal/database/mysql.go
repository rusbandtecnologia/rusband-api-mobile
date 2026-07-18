package database

import (
	"fmt"
	"log"

	"github.com/rusbandtecnologia/rusband-api-mobile/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(app config.App) {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		app.DBUsername,
		app.DBPassword,
		app.DBHost,
		app.DBPort,
		app.DBDatabase,
	)

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados.")
	}

	DB = connection

	log.Println("Banco de dados conectado.")

}

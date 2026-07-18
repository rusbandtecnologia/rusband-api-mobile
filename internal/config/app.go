package config

import "os"

type App struct {
	Name string
	Port string

	DBHost     string
	DBPort     string
	DBDatabase string
	DBUsername string
	DBPassword string

	JWTSecret string
}

func LoadApp() App {

	return App{
		Name: os.Getenv("APP_NAME"),
		Port: os.Getenv("APP_PORT"),

		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBDatabase: os.Getenv("DB_DATABASE"),
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),

		JWTSecret: os.Getenv("JWT_SECRET"),
	}

}

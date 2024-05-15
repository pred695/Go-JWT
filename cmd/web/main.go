package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/pred695/Go-JWT/Routes"
	"github.com/pred695/Go-JWT/database"
	"gorm.io/gorm"
)

var DbConn *gorm.DB
var Config database.Config

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Config = database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}
	DbConn, err = database.Connect(&Config)
}
func main() {
	app := fiber.New()
	db, err := DbConn.DB()
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	defer db.Close()
	Routes.SetUpTaskRoutes(app)

	Routes.SetUpUserRoutes(app)
	app.Listen(":3000")
}

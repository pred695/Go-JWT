package database

import (
	"fmt"

	"github.com/pred695/Go-JWT/Models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

var DbConn *gorm.DB

func Connect(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),	//logging the queries
	})
	if err != nil {
		fmt.Println("Error connecting to database")
		return db, err
	}

	fmt.Println("Connection successful")
	db.AutoMigrate(new(Models.Task)) //creating the table if not exists
	db.AutoMigrate(new(Models.User)) //creating the table if not exists
	DbConn = db
	return db, nil
}

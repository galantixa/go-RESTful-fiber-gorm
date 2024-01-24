package database

import (
	"fmt"
	"github.com/galantixa/gofiber-gorm/config"
	"github.com/galantixa/gofiber-gorm/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"strconv"
)

// declare variable for the database
var DB *gorm.DB

// connect DB

func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseInt(p, 10, 32)

	if err != nil {
		log.Println("LOL")
	}
	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	// connect db and innitialize db connect varibale
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("Failed to connect database!")
	}
	fmt.Println("Connection Opened to Database")

	// migrate the databases
	DB.AutoMigrate(&model.Note{})
	fmt.Println("DB migrated")
}

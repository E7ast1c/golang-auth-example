package dao

import (
	"fmt"
	"github.com/E7ast1c/golang-auth-example/api/models"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func ConnectDB() *gorm.DB {

	//Load environmenatal variables
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("databaseUser")
	password := os.Getenv("databasePassword")
	databaseName := os.Getenv("databaseName")
	databaseHost := os.Getenv("databaseHost")

	//Define DB connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", databaseHost, username, databaseName, password)

	//connect to db URI
	db, err := gorm.Open("postgres", dbURI)

	if err != nil {
		fmt.Println("error", err)
		panic(err)
	}
	// close db when not in use
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(
		&models.User{})

	fmt.Println("Successfully connected!", db)
	return db
}
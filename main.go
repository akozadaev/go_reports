package main

import (
	database "akozadaev/go_reports/db"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

func init() {
	envPath := "."
	envFileName := ".env"

	fullPath := envPath + "/" + envFileName

	if err := godotenv.Overload(fullPath); err != nil {
		log.Printf("[ERROR] failed with %+v", "No .env file found")
	}
}

func main() {
	dbName := os.Getenv("DB_NAME")
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	database.Migrate(db)

}

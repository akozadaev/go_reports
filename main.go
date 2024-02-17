package main

import (
	"github.com/joho/godotenv"
	"log"
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

}

package common

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	// Check if .env.local file exists
	envFile := ".env.local"
	if _, err := os.Stat(envFile); os.IsNotExist(err) {
		log.Default().Println("No .env.local file found, finding .env.prod")
		envFile = ".env.qa"
	}

	if _, err := os.Stat(envFile); os.IsNotExist(err) {
		log.Default().Println("No .env.qa file found, using .env.prod")
		envFile = ".env.prod"
	}

	// Load environment variables from the chosen .env file
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading %s file", envFile)
	}
	log.Default().Println("Currenly using " + envFile + " environment")
}

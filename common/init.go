package common

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	// Check if .env.local file exists
	envFile := ".env.local"
	if _, err := os.Stat(envFile); os.IsNotExist(err) {
		log.Default().Println("No .env.local file found, finding .env")
		envFile = ".env"
	}

	// Load environment variables from the chosen .env file
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading %s file", envFile)
	}
	log.Default().Println("Currenly using " + envFile + " environment")
}

var loc *time.Location

func SetTimezone(tz string) error {
	location, err := time.LoadLocation(tz)
	if err != nil {
		return err
	}
	loc = location
	return nil
}

func GetTime(t time.Time) time.Time {
	return t.In(loc)
}

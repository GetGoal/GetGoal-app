package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type (
	Config struct {
		App               App
		Db                Db
		Env               string
		Search            Search
		EnvFromDockerFile DockerEnv
	}

	App struct {
		Port int
	}

	Db struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
		SSLMode  string
		TimeZone string
	}
	Search struct {
		LabelLimit int
	}

	DockerEnv struct {
		Env      string
		DbHost   string
		DbPort   int
		DbUser   string
		DbPass   string
		DbName   string
		TimeZone string
	}
)

func GetConfig() Config {
	// Set up Viper
	viper.AutomaticEnv() // Automatically read from environment variables
	viper.SetConfigType("env")

	log.Default().Println("Binding environment variables...")
	viper.BindEnv("ENV")
	viper.BindEnv("DB_HOST")
	viper.BindEnv("DB_USER")
	viper.BindEnv("DB_PASSWORD")
	viper.BindEnv("DB_NAME")
	viper.BindEnv("DB_PORT")
	viper.BindEnv("TZ")
	log.Default().Println("Done binding environment variables")

	log.Default().Println("Reading config file...")
	viper.SetConfigName("config.local")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		log.Default().Printf("Fatal error when loading config.local file: %s \n", err)

		log.Default().Printf("Reading config.yaml")
		viper.SetConfigName("config")

		if err := viper.ReadInConfig(); err != nil {
			log.Default().Panicf("Fatal error when loading config file: %s \n", err)
		}
	}

	log.Default().Printf("Using config file: %s \n", strings.Split(viper.ConfigFileUsed(), "/")[len(strings.Split(viper.ConfigFileUsed(), "/"))-1])

	return Config{
		App: App{
			Port: viper.GetInt("app.server.port"),
		},
		Db: Db{
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetInt("DB_PORT"),
			User:     viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
			DBName:   viper.GetString("DB_NAME"),
			SSLMode:  "disable",
			TimeZone: viper.GetString("TZ"),
		},
		Env: viper.GetString("env"),
		Search: Search{
			LabelLimit: viper.GetInt("search.label_limit"),
		},
	}
}

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

	viper.BindEnv("ENV")
	viper.BindEnv("DB_HOST")
	viper.BindEnv("DB_USER")
	viper.BindEnv("DB_PASSWORD")
	viper.BindEnv("DB_NAME")
	viper.BindEnv("DB_PORT")
	viper.BindEnv("TZ")
	log.Default().Printf(" ENV variables in dockerfiles :")
	log.Default().Printf(" ENV : %s", viper.GetString("ENV"))
	log.Default().Printf(" DB_HOST : %s", viper.GetString("DB_HOST"))
	log.Default().Printf(" DB_USER : %s", viper.GetString("DB_USER"))
	log.Default().Printf(" DB_PASSWORD : %s", viper.GetString("DB_PASSWORD"))
	log.Default().Printf(" DB_NAME : %s", viper.GetString("DB_NAME"))
	log.Default().Printf(" DB_PORT : %s", viper.GetString("DB_PORT"))
	log.Default().Printf(" TZ : %s", viper.GetString("TZ"))

	log.Default().Println("Reading config file...")
	viper.SetConfigName("config.local")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		log.Default().Printf("Fatal error when loading config.local file: %s \n", err)
		log.Default().Printf("Using config.yaml")

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
			Host:     viper.GetString("database.host"),
			Port:     viper.GetInt("database.port"),
			User:     viper.GetString("database.user"),
			Password: viper.GetString("database.password"),
			DBName:   viper.GetString("database.dbname"),
			SSLMode:  viper.GetString("database.sslmode"),
			TimeZone: viper.GetString("database.timezone"),
		},
		Env: viper.GetString("env"),
		Search: Search{
			LabelLimit: viper.GetInt("search.label_limit"),
		},
	}
}

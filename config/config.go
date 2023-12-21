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

	log.Default().Println("Binding environment variables...")
	viper.BindEnv("env")
	viper.BindEnv("database.host")
	viper.BindEnv("database.port")
	viper.BindEnv("database.user")
	viper.BindEnv("database.password")
	viper.BindEnv("database.dbname")
	viper.BindEnv("database.sslmode")
	viper.BindEnv("database.timezone")

	viper.BindEnv("ENV")
	viper.BindEnv("DB_HOST")
	viper.BindEnv("DB_USER")
	viper.BindEnv("DB_PASSWORD")
	viper.BindEnv("DB_NAME")
	viper.BindEnv("DB_PORT")
	viper.BindEnv("TZ")

	log.Default().Printf("ENV in docker container : ")
	log.Default().Printf("Env %s", viper.GetString("env"))
	log.Default().Printf("DbHost %s", viper.GetString("database.host"))
	log.Default().Printf("DbPort %d", viper.GetInt("database.port"))
	log.Default().Printf("DbUser %s", viper.GetString("database.user"))
	log.Default().Printf("DbPass %s", viper.GetString("database.password"))
	log.Default().Printf("DbName %s", viper.GetString("database.dbname"))
	log.Default().Printf("SSLMode %s", viper.GetString("database.sslmode"))
	log.Default().Printf("TimeZone %s", viper.GetString("database.timezone"))

	log.Default().Printf("ENV in docker container : ")
	log.Default().Printf("Env %s", viper.GetString("ENV"))
	log.Default().Printf("DbHost %s", viper.GetString("DB_HOST"))
	log.Default().Printf("DbPort %d", viper.GetInt("DB_PORT"))
	log.Default().Printf("DbUser %s", viper.GetString("DB_USER"))
	log.Default().Printf("DbPass %s", viper.GetString("DB_PASSWORD"))
	log.Default().Printf("DbName %s", viper.GetString("DB_NAME"))
	log.Default().Printf("TimeZone %s", viper.GetString("TZ"))

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

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

	viper.Set("env", viper.BindEnv("env"))
	viper.Set("database.host", viper.BindEnv("database_host"))
	viper.Set("database.port", viper.BindEnv("database_port"))
	viper.Set("database.user", viper.BindEnv("database_user"))
	viper.Set("database.password", viper.BindEnv("database_password"))
	viper.Set("database.dbname", viper.BindEnv("database_dbname"))
	viper.Set("database.sslmode", viper.BindEnv("database_sslmode"))
	viper.Set("database.timezone", viper.BindEnv("database_timezone"))
	log.Default().Println(viper.AllSettings())

	viper.BindEnv("env")
	viper.BindEnv("database_host")
	viper.BindEnv("database_port")
	viper.BindEnv("database_user")
	viper.BindEnv("database_password")
	viper.BindEnv("database_dbname")
	viper.BindEnv("database_sslmode")
	viper.BindEnv("database_timezone")
	// viper.BindEnv("database.host")
	// viper.BindEnv("database.port")
	// viper.BindEnv("database.user")
	// viper.BindEnv("database.password")
	// viper.BindEnv("database.dbname")
	// viper.BindEnv("database.sslmode")
	// viper.BindEnv("database.timezone")

	log.Default().Printf("ENV in docker container : ")
	log.Default().Printf("Env %s", viper.GetString("env"))
	log.Default().Printf("DbHost %s", viper.GetString("database.host"))
	log.Default().Printf("DbPort %d", viper.GetInt("database.port"))
	log.Default().Printf("DbUser %s", viper.GetString("database.user"))
	log.Default().Printf("DbPass %s", viper.GetString("database.password"))
	log.Default().Printf("DbName %s", viper.GetString("database.dbname"))
	log.Default().Printf("SSLMode %s", viper.GetString("database.sslmode"))
	log.Default().Printf("TimeZone %s", viper.GetString("database.timezone"))

	log.Default().Println("Done binding environment variables")
	log.Default().Println(viper.AllSettings())

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

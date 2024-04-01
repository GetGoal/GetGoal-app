package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

var config Config

type (
	Config struct {
		App               App
		Db                Db
		Env               string
		Search            Search
		EnvFromDockerFile DockerEnv
		Mailer            MailerConfig
		JwtKeys           JwtKey
		DevGorseConfig    GorseConfig
		QaGorseConfig     GorseConfig
		ProdGorseConfig   GorseConfig
		Recommendation    Recommendation
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
		LabelLimit      int
		PreferenceLimit int
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

	MailerConfig struct {
		Host     string
		Port     int
		Email    string
		Password string
		BaseURL  string
	}
	JwtKey struct {
		AccessSecret  string
		RefreshSecret string
	}

	GorseConfig struct {
		Host string
		Port int
	}

	Recommendation struct {
		Limit int
	}
)

func ReadConfig(path string) Config {

	log.Default().Println("Reading environment variables...")
	viper.AutomaticEnv() // Automatically read from environment variables
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	log.Default().Println("Reading config file...")
	viper.SetConfigName("config.local")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		log.Default().Printf("Fatal error when loading config.local file: %s \n", err)

		log.Default().Printf("Reading config.yaml")
		viper.SetConfigName("config")

		if err := viper.ReadInConfig(); err != nil {
			log.Default().Panicf("Fatal error when loading config file: %s \n", err)
		}
	}

	log.Default().Printf("Using config file: %s \n", strings.Split(viper.ConfigFileUsed(), "/")[len(strings.Split(viper.ConfigFileUsed(), "/"))-1])

	config = Config{
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
			LabelLimit:      viper.GetInt("search.label_limit"),
			PreferenceLimit: viper.GetInt("search.preference_limit"),
		},
		Mailer: MailerConfig{
			Host:     viper.GetString("mailer.host"),
			Port:     viper.GetInt("mailer.port"),
			Email:    viper.GetString("mailer.email"),
			Password: viper.GetString("mailer.password"),
			BaseURL:  viper.GetString("mailer.url"),
		},
		JwtKeys: JwtKey{
			AccessSecret:  viper.GetString("secrets.jwt.accesskey"),
			RefreshSecret: viper.GetString("secrets.jwt.refreshkey"),
		},
		DevGorseConfig: GorseConfig{
			Host: viper.GetString("gorse.dev.host"),
			Port: viper.GetInt("gorse.dev.port"),
		},
		QaGorseConfig: GorseConfig{
			Host: viper.GetString("gorse.qa.host"),
			Port: viper.GetInt("gorse.qa.port"),
		},
		ProdGorseConfig: GorseConfig{
			Host: viper.GetString("gorse.prod.host"),
			Port: viper.GetInt("gorse.prod.port"),
		},
		Recommendation: Recommendation{
			Limit: viper.GetInt("recommendation.limit"),
		},
	}
	return config
}

func GetConfig() Config {
	return config
}

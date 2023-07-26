package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type AppConfig struct {
	JWT    string
	DBUSER string
	DBPASS string
	DBHOST string
	DBPORT string
	DBNAME string
}

func InitConfig() *AppConfig {
	return readEnv()
}

func readEnv() *AppConfig {
	app := AppConfig{}

	if !loadFromEnv(&app) {
		if !loadFromConfigFile(&app) {
			log.Println("error reading config")
			return nil
		}
	}

	return &app
}

func loadFromEnv(app *AppConfig) bool {
	isRead := false

	if val, found := os.LookupEnv("DBUSER"); found {
		app.DBUSER = val
		isRead = true
	}

	if val, found := os.LookupEnv("DBPASS"); found {
		app.DBPASS = val
		isRead = true
	}

	if val, found := os.LookupEnv("DBHOST"); found {
		app.DBHOST = val
		isRead = true
	}

	if val, found := os.LookupEnv("DBPORT"); found {
		app.DBPORT = val
		isRead = true
	}

	if val, found := os.LookupEnv("DBNAME"); found {
		app.DBNAME = val
		isRead = true
	}

	if val, found := os.LookupEnv("JWT"); found {
		app.JWT = val
		isRead = true
	}

	return isRead
}

func loadFromConfigFile(app *AppConfig) bool {
	viper.AddConfigPath(".")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("error reading config:", err.Error())
		return false
	}

	app.JWT = viper.GetString("JWT")
	app.DBUSER = viper.GetString("DBUSER")
	app.DBPASS = viper.GetString("DBPASS")
	app.DBHOST = viper.GetString("DBHOST")
	app.DBPORT = viper.GetString("DBPORT")
	app.DBNAME = viper.GetString("DBNAME")

	return true
}

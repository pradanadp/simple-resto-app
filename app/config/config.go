package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var (
	JWT string
)

type AppConfig struct {
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
	isRead := true

	if val, found := os.LookupEnv("DBUSER"); found {
		app.DBUSER = val
		isRead = false
	}

	if val, found := os.LookupEnv("DBPASS"); found {
		app.DBPASS = val
		isRead = false
	}

	if val, found := os.LookupEnv("DBHOST"); found {
		app.DBHOST = val
		isRead = false
	}

	if val, found := os.LookupEnv("DBPORT"); found {
		app.DBPORT = val
		isRead = false
	}

	if val, found := os.LookupEnv("DBNAME"); found {
		app.DBNAME = val
		isRead = false
	}

	if val, found := os.LookupEnv("JWT"); found {
		JWT = val
		isRead = false
	}

	if isRead {
		viper.AddConfigPath(".")
		viper.SetConfigName("local")
		viper.SetConfigType("yaml")

		err := viper.ReadInConfig()
		if err != nil {
			log.Println("error read config : ", err.Error())
			return nil
		}

		JWT = viper.GetString("JWT")
		app.DBUSER = viper.GetString("DBUSER")
		app.DBPASS = viper.GetString("DBPASS")
		app.DBHOST = viper.GetString("DBHOST")
		app.DBPORT = viper.GetString("DBPORT")
		app.DBNAME = viper.GetString("DBNAME")
	}

	return &app
}

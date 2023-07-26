package database

import (
	"fmt"

	"github.com/pradanadp/simple-resto-app/app/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	log = logrus.New()
)

func InitDB(c *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		c.DBHOST, c.DBUSER, c.DBPASS, c.DBNAME, c.DBPORT,
	)

	db := openDB(dsn)

	// err = db.AutoMigrate(
	// 	&user.User{},
	// 	&user.Position{},
	// 	&company.Company{},
	// 	&transaction.Transaction{},
	// )
	// if err != nil {
	// 	log.Panic(err.Error())
	// 	panic(err.Error())
	// }

	log.Info("Success to connect and migrate to database")
	return db
}

func openDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err.Error())
		panic(err.Error())
	}

	return db
}

// func migrateDB()

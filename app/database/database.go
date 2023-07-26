package database

import (
	"fmt"

	"github.com/pradanadp/simple-resto-app/app/config"
	_customer "github.com/pradanadp/simple-resto-app/features/customer/repository"
	_order "github.com/pradanadp/simple-resto-app/features/order/repository"
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

	err := db.AutoMigrate(
		&_customer.Customer{},
		&_order.Order{},
		&_order.Cart{},
		&_order.Item{},
		&_order.OrderItem{},
		&_order.PurchaseReceipt{},
	)
	if err != nil {
		log.Panic(err.Error())
		panic(err.Error())
	}

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

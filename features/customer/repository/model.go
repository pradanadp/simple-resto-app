package repository

import (
	"time"

	_order "github.com/pradanadp/simple-resto-app/features/order/repository"
	"gorm.io/gorm"
)

type Customer struct {
	CustomerID  string         `gorm:"primaryKey;type:varchar(255)"`
	FirstName   string         `gorm:"type:varchar(255);not null"`
	LastName    string         `gorm:"type:varchar(255);not null"`
	PhoneNumber string         `gorm:"type:varchar(255);not null"`
	Email       string         `gorm:"type:varchar(255);not null"`
	Address     string         `gorm:"type:text;not null"`
	Orders      []_order.Order `gorm:"foreignKey:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt   time.Time      `gorm:"type:datetime"`
	UpdatedAt   time.Time      `gorm:"type:datetime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

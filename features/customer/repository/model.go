package repository

import (
	"time"

	"github.com/pradanadp/simple-resto-app/features/customer"
	"gorm.io/gorm"
)

type Customer struct {
	CustomerID  string         `gorm:"primaryKey;type:varchar(255)"`
	FirstName   string         `gorm:"type:varchar(255);not null"`
	LastName    string         `gorm:"type:varchar(255);not null"`
	PhoneNumber string         `gorm:"type:varchar(255);not null"`
	Email       string         `gorm:"type:varchar(255);not null"`
	Password    string         `gorm:"type:varchar(255);not null"`
	Address     string         `gorm:"type:text;not null"`
	Orders      []Order        `gorm:"foreignKey:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt   time.Time      `gorm:"type:datetime"`
	UpdatedAt   time.Time      `gorm:"type:datetime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Order struct {
	OrderID       string         `gorm:"primaryKey;type:varchar(255)"`
	CustomerID    string         `gorm:"type:varchar(255);not null"`
	TotalAmount   float64        `gorm:"type:decimal(10,2);not null"`
	PaymentStatus string         `gorm:"type:('paid', 'pending', 'cancelled');default:'pending'"`
	OrderStatus   string         `gorm:"type:('pending', 'preparing', 'ready', 'delivery');default:'pending'"`
	DeliveryAddr  string         `gorm:"type:text;not null"`
	ContactInfo   string         `gorm:"type:varchar(255);not null"`
	CreatedAt     time.Time      `gorm:"type:datetime"`
	UpdatedAt     time.Time      `gorm:"type:datetime"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

func CustomerEntityToModel(c customer.CustomerEntity) Customer {
	return Customer{
		CustomerID:  c.CustomerID,
		FirstName:   c.FirstName,
		LastName:    c.LastName,
		PhoneNumber: c.PhoneNumber,
		Email:       c.Email,
		Password:    c.Password,
		Address:     c.Address,
	}
}

func CustomerModelToEntity(c Customer) customer.CustomerEntity {
	return customer.CustomerEntity{
		CustomerID:  c.CustomerID,
		FirstName:   c.FirstName,
		LastName:    c.LastName,
		PhoneNumber: c.PhoneNumber,
		Email:       c.Email,
		Address:     c.Address,
	}
}

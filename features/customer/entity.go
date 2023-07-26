package customer

import (
	"time"
)

type CustomerEntity struct {
	CustomerID  string
	FirstName   string
	LastName    string
	PhoneNumber string
	Email       string
	Address     string
	Orders      []OrderEntity
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type OrderEntity struct {
	OrderID       string
	CustomerID    string
	TotalAmount   float64
	PaymentStatus string
	OrderStatus   string
	DeliveryAddr  string
	ContactInfo   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

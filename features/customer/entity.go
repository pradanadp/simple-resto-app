package customer

import (
	"time"

	"github.com/labstack/echo/v4"
)

type CustomerEntity struct {
	CustomerID  string
	FirstName   string
	LastName    string
	PhoneNumber string
	Email       string
	Password    string
	Address     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	Orders      []OrderEntity
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

type Repository interface {
	Register(req CustomerEntity) (CustomerEntity, error)
	Login(req CustomerEntity) (CustomerEntity, string, error)
}

type Service interface {
	Register(req CustomerEntity) (CustomerEntity, error)
	Login(req CustomerEntity) (CustomerEntity, string, error)
}

type Controller interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
}

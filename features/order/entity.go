package order

import (
	"time"

	"github.com/labstack/echo/v4"
)

type OrderEntity struct {
	OrderID       string
	CustomerID    string
	TotalAmount   float64
	PaymentStatus string
	OrderStatus   string
	DeliveryAddr  string
	ContactInfo   string
	OrderItems    []OrderItemEntity
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

type OrderItemEntity struct {
	OrderItemID string
	MenuItemID  string
	Quantity    uint
	Subtotal    float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type PurchaseReceiptEntity struct {
	PurchaseReceiptID string
	PaymentMethod     string
	ReceiptTotal      float64
	ReceiptDate       time.Time
	AdditionalDetails string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         time.Time
}

type CartEntity struct {
	CartID     string
	CustomerID string
	ItemID     string
	Quantity   uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

type Repository interface {
	GetReceipt(orderID string) (PurchaseReceiptEntity, error)
	AddToCart(itemID string)
	RemoveFromCart(itemID string)
	GetWeeklyIncomeReport(start_date, end_date time.Time)
	GetMonthlyIncomeReport()
}

type Service interface {
	GetReceipt(orderID string) (PurchaseReceiptEntity, error)
	AddToCart(itemID string)
	RemoveFromCart(itemID string)
	GetWeeklyIncomeReport(start_date, end_date time.Time)
	GetMonthlyIncomeReport()
}

type Controller interface {
	GetReceipt() echo.HandlerFunc
	AddToCart() echo.HandlerFunc
	GetIncomeReport() echo.HandlerFunc
}

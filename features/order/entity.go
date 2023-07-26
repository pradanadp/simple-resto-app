package order

import (
	"time"
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

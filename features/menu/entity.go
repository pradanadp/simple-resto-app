package order

import (
	"time"
)

type MenuItemEntity struct {
	MenuItemID   string
	ItemName     string
	Description  string
	Price        float64
	Category     string
	Availability string
	OrderItems   []OrderItemEntity
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
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

package repository

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	OrderID       string         `gorm:"primaryKey;type:varchar(255)"`
	CustomerID    string         `gorm:"type:varchar(255);not null"`
	TotalAmount   float64        `gorm:"type:decimal(10,2);not null"`
	PaymentStatus string         `gorm:"type:('paid', 'pending', 'cancelled');default:'pending'"`
	OrderStatus   string         `gorm:"type:('pending', 'preparing', 'ready', 'delivery');default:'pending'"`
	DeliveryAddr  string         `gorm:"type:text;not null"`
	ContactInfo   string         `gorm:"type:varchar(255);not null"`
	OrderItems    []OrderItem    `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt     time.Time      `gorm:"type:datetime"`
	UpdatedAt     time.Time      `gorm:"type:datetime"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type OrderItem struct {
	OrderItemID string         `gorm:"primaryKey;type:varchar(255)"`
	OrderID     string         `gorm:"type:varchar(255);not null"`
	MenuItemID  string         `gorm:"type:varchar(255);not null"`
	Quantity    uint           `gorm:"type:integer;not null"`
	Subtotal    float64        `gorm:"type:decimal(10,2);not null"`
	CreatedAt   time.Time      `gorm:"type:datetime"`
	UpdatedAt   time.Time      `gorm:"type:datetime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type PurchaseReceipt struct {
	PurchaseReceiptID string         `gorm:"primaryKey;type:varchar(255)"`
	PaymentMethod     string         `gorm:"type:('credit_card', 'cash', 'transfer', 'e_wallet', 'e_money', 'qr_code');default:'cash'"`
	ReceiptTotal      float64        `gorm:"type:decimal(10,2);not null"`
	ReceiptDate       time.Time      `gorm:"type:datetime"`
	AdditionalDetails string         `gorm:"type:text"`
	CreatedAt         time.Time      `gorm:"type:datetime"`
	UpdatedAt         time.Time      `gorm:"type:datetime"`
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

package repository

import (
	"time"

	"github.com/pradanadp/simple-resto-app/features/order"
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
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

type OrderItem struct {
	OrderItemID string         `gorm:"primaryKey;type:varchar(255)"`
	OrderID     string         `gorm:"type:varchar(255);not null"`
	ItemID      string         `gorm:"type:varchar(255);not null"`
	Quantity    uint           `gorm:"type:integer;not null"`
	Subtotal    float64        `gorm:"type:decimal(10,2);not null"`
	CreatedAt   time.Time      `gorm:"type:datetime"`
	UpdatedAt   time.Time      `gorm:"type:datetime"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type Cart struct {
	CartID     string `gorm:"primaryKey;type:varchar(255)"`
	CustomerID string `gorm:"type:varchar(255);not null"`
	Customer   Customer
	ItemID     string         `gorm:"type:varchar(255);not null"`
	Quantity   uint           `gorm:"type:integer;not null"`
	Items      []Item         `gorm:"foreignKey:ItemID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt  time.Time      `gorm:"type:datetime"`
	UpdatedAt  time.Time      `gorm:"type:datetime"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

type Customer struct {
	CustomerID  string         `gorm:"primaryKey;type:varchar(255)"`
	FirstName   string         `gorm:"type:varchar(255);not null"`
	LastName    string         `gorm:"type:varchar(255);not null"`
	PhoneNumber string         `gorm:"type:varchar(255);not null"`
	Email       string         `gorm:"type:varchar(255);not null"`
	Password    string         `gorm:"type:varchar(255);not null"`
	Address     string         `gorm:"type:text;not null"`
	CreatedAt   time.Time      `gorm:"type:datetime"`
	UpdatedAt   time.Time      `gorm:"type:datetime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Item struct {
	CartID      string
	ItemID      string         `gorm:"primaryKey;type:varchar(255)"`
	ItemName    string         `gorm:"type:varchar(255);not null"`
	Description string         `gorm:"type:text;not null"`
	Price       float64        `gorm:"type:decimal(10,2);not null"`
	Category    string         `gorm:"type:('appetizers', 'main_course', 'desserts');default:'main_course'"`
	Quantity    uint           `gorm:"type:integer"`
	OrderItems  []OrderItem    `gorm:"foreignKey:MenuItemID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
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
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
}

func CartEntityToModel(c order.CartEntity) Cart {
	return Cart{
		CustomerID: c.CustomerID,
		ItemID:     c.ItemID,
		Quantity:   c.Quantity,
	}
}

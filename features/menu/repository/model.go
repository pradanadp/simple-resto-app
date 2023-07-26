package repository

import (
	"time"

	_orderItem "github.com/pradanadp/simple-resto-app/features/order/repository"
	"gorm.io/gorm"
)

type MenuItem struct {
	MenuItemID   string                 `gorm:"primaryKey;type:varchar(255)"`
	ItemName     string                 `gorm:"type:varchar(255);not null"`
	Description  string                 `gorm:"type:text;not null"`
	Price        float64                `gorm:"type:decimal(10,2);not null"`
	Category     string                 `gorm:"type:('appetizers', 'main_course', 'desserts');default:'main_course'"`
	Availability string                 `gorm:"type:('in_stock', 'out_of_stock');default:'in_stock'"`
	OrderItems   []_orderItem.OrderItem `gorm:"foreignKey:MenuItemID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt    time.Time              `gorm:"type:datetime"`
	UpdatedAt    time.Time              `gorm:"type:datetime"`
	DeletedAt    gorm.DeletedAt         `gorm:"index"`
}

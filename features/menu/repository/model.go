package repository

import (
	"time"

	"github.com/pradanadp/simple-resto-app/features/menu"

	"gorm.io/gorm"
)

type Item struct {
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

type OrderItem struct {
	OrderItemID string         `gorm:"primaryKey;type:varchar(255)"`
	OrderID     string         `gorm:"type:varchar(255);not null"`
	ItemID      string         `gorm:"type:varchar(255);not null"`
	Quantity    uint           `gorm:"type:integer;not null"`
	Subtotal    float64        `gorm:"type:decimal(10,2);not null"`
	CreatedAt   time.Time      `gorm:"type:datetime"`
	UpdatedAt   time.Time      `gorm:"type:datetime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func ItemEntityToModel(i menu.ItemEntity) Item {
	return Item{
		ItemID:      i.ItemID,
		ItemName:    i.ItemName,
		Description: i.Description,
		Price:       i.Price,
		Category:    i.Category,
		Quantity:    i.Quantity,
	}
}

func ItemModelToEntity(i Item) menu.ItemEntity {
	return menu.ItemEntity{
		ItemID:      i.ItemID,
		ItemName:    i.ItemName,
		Description: i.Description,
		Price:       i.Price,
		Category:    i.Category,
		Quantity:    i.Quantity,
		CreatedAt:   i.CreatedAt,
		UpdatedAt:   i.UpdatedAt,
		DeletedAt:   i.DeletedAt.Time,
	}
}

func OrderItemModelToEntity(o OrderItem) menu.OrderItemEntity {
	return menu.OrderItemEntity{
		OrderItemID: o.OrderItemID,
		OrderID:     o.OrderID,
		ItemID:      o.ItemID,
		Quantity:    o.Quantity,
		Subtotal:    o.Subtotal,
		CreatedAt:   o.CreatedAt,
		UpdatedAt:   o.UpdatedAt,
		DeletedAt:   o.DeletedAt.Time,
	}
}

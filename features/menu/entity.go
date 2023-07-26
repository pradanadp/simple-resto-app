package menu

import (
	"time"

	"github.com/labstack/echo/v4"
)

type ItemEntity struct {
	ItemID      string
	ItemName    string
	Description string
	Price       float64
	Category    string
	Quantity    uint
	OrderItems  []OrderItemEntity
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type OrderItemEntity struct {
	OrderItemID string
	OrderID     string
	ItemID      string
	Quantity    uint
	Subtotal    float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type StockReportEntity struct {
	ItemID            string
	ReportDate        time.Time
	AvailableQuantity uint
	TotalSoldQuantity uint
}

type Repository interface {
	GetStockReport(itemID string) (StockReportEntity, error)
}

type Service interface {
	GetStockReport(itemID string) (StockReportEntity, error)
}

type Controller interface {
	GetStockReport() echo.HandlerFunc
}

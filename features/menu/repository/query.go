package repository

import (
	"time"

	"github.com/pradanadp/simple-resto-app/features/menu"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	log = logrus.New()
)

type menuQuery struct {
	db *gorm.DB
}

// GetStockReport implements menu.MenuRepository.
func (mq *menuQuery) GetStockReport(itemID string) (menu.StockReportEntity, error) {
	stockReport := menu.StockReportEntity{}
	stockReport.ItemID = itemID
	stockReport.ReportDate = time.Now()

	query := `
		SELECT
			item.quantity AS available_quantity,
			SUM(order_items.quantity) AS total_sold_quantity
		FROM
			item
		LEFT JOIN
			order_items ON item.item_id = order_items.item_id
		WHERE
			item.item_id = ?
		GROUP BY
			item.quantity
	`
	err := mq.db.Raw(query, itemID).Scan(&stockReport).Error
	if err != nil {
		log.Error(err)
		return menu.StockReportEntity{}, err
	}

	return stockReport, nil
}

func New(db *gorm.DB) menu.Repository {
	return &menuQuery{
		db: db,
	}
}

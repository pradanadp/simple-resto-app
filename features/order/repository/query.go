package repository

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/pradanadp/simple-resto-app/features/order"
	"github.com/pradanadp/simple-resto-app/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	log = logrus.New()
)

type orderQuery struct {
	db *gorm.DB
}

// AddItemToCart implements order.Repository.
func (oq *orderQuery) AddItemToCart(itemID, customerID string, quantity uint) error {
	tx := oq.db.Begin()
	cartID := utils.GenerateCartID()

	cart := Cart{
		CartID:     cartID,
		CustomerID: customerID,
		ItemID:     itemID,
		Quantity:   quantity,
	}

	if err := tx.Create(&cart).Error; err != nil {
		tx.Rollback()
		return err
	}

	if tx.RowsAffected == 0 {
		tx.Rollback()
		log.Error("no row affected. failed to add item to cart")
		return errors.New("failed to insert, row affected is 0")
	}

	return tx.Commit().Error
}

// GetMonthlyIncomeReport implements order.Repository.
func (oq *orderQuery) GetMonthlyIncomeReport(yearMonth string) (order.IncomeReportEntity, error) {
	incomeReport := order.IncomeReportEntity{}

	query := `
		SELECT SUM(total_amount) AS monthly_total_income
		FROM orders
		WHERE EXTRACT(YEAR_MONTH FROM order_date) = ?
	`

	var totalIncome float64
	if err := oq.db.Raw(query, yearMonth).Scan(&totalIncome).Error; err != nil {
		return incomeReport, err
	}

	year, err := strconv.Atoi(yearMonth[:4])
	if err != nil {
		return incomeReport, fmt.Errorf("failed to parse year from year_month: %v", err)
	}

	month, err := strconv.Atoi(yearMonth[4:])
	if err != nil {
		return incomeReport, fmt.Errorf("failed to parse month from year_month: %v", err)
	}

	// Create the time.Time object representing the start of the month
	startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)

	// Determine the end of the month
	endDate := startDate.AddDate(0, 1, 0).Add(-time.Nanosecond)

	incomeReport.StartDate = startDate
	incomeReport.EndDate = endDate
	incomeReport.TotalIncome = totalIncome

	return incomeReport, nil
}

// GetReceipt implements order.Repository.
func (oq *orderQuery) GetReceipt(orderID string) (order.PurchaseReceiptEntity, error) {
	receipt := order.PurchaseReceiptEntity{}

	query := oq.db.Table("purchase_receipts").Where("order_id = ?", orderID).First(&receipt)
	if query.Error != nil {
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			log.Error(errors.New("purchase receipt not found"))
			return receipt, errors.New("purchase receipt not found")
		}
		return receipt, query.Error
	}

	return receipt, nil
}

// GetWeeklyIncomeReport implements order.Repository.
func (oq *orderQuery) GetWeeklyIncomeReport(startDate, endDate string) (order.IncomeReportEntity, error) {
	incomeReport := order.IncomeReportEntity{}

	query := `
		SELECT SUM(total_amount) AS weekly_total_income
		FROM orders
		WHERE order_date >= ? AND order_date <= ?
	`

	startDateObj, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		log.Error(err)
		return incomeReport, fmt.Errorf("failed to parse start date: %v", err)
	}

	endDateObj, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		log.Error(err)
		return incomeReport, fmt.Errorf("failed to parse end date: %v", err)
	}

	var totalIncome float64
	if err := oq.db.Raw(query, startDateObj, endDateObj).Scan(&totalIncome).Error; err != nil {
		return incomeReport, err
	}

	incomeReport.StartDate = startDateObj
	incomeReport.EndDate = endDateObj
	incomeReport.TotalIncome = totalIncome

	return incomeReport, nil
}

// RemoveItemFromCart implements order.Repository.
func (oq *orderQuery) RemoveItemFromCart(itemID string) error {
	var cart Cart
	if err := oq.db.Where("item_id = ?", itemID).First(&cart).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("item not found in the cart")
		}
		return err
	}

	tx := oq.db.Begin()

	if err := tx.Delete(&cart).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func New(db *gorm.DB) order.Repository {
	return &orderQuery{
		db: db,
	}
}

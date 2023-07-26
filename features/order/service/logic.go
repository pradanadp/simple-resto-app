package service

import (
	"github.com/pradanadp/simple-resto-app/features/order"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.New()
)

type orderService struct {
	repository order.Repository
}

// AddItemToCart implements order.Service.
func (os *orderService) AddItemToCart(itemID, customerID string, quantity uint) error {
	err := os.repository.AddItemToCart(itemID, customerID, quantity)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// GetMonthlyIncomeReport implements order.Service.
func (os *orderService) GetMonthlyIncomeReport(yearMonth string) (order.IncomeReportEntity, error) {
	incomeReport, err := os.repository.GetMonthlyIncomeReport(yearMonth)
	if err != nil {
		return order.IncomeReportEntity{}, err
	}

	return incomeReport, nil
}

// GetReceipt implements order.Service.
func (os *orderService) GetReceipt(orderID string) (order.PurchaseReceiptEntity, error) {
	receipt, err := os.repository.GetReceipt(orderID)
	if err != nil {
		return order.PurchaseReceiptEntity{}, err
	}

	return receipt, nil
}

// GetWeeklyIncomeReport implements order.Service.
func (os *orderService) GetWeeklyIncomeReport(startDate, endDate string) (order.IncomeReportEntity, error) {
	incomeReport, err := os.repository.GetWeeklyIncomeReport(startDate, endDate)
	if err != nil {
		return order.IncomeReportEntity{}, err
	}

	return incomeReport, nil
}

// RemoveItemFromCart implements order.Service.
func (os *orderService) RemoveItemFromCart(itemID string) error {
	err := os.repository.RemoveItemFromCart(itemID)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func New(r order.Repository) order.Service {
	return &orderService{
		repository: r,
	}
}

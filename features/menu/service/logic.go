package service

import (
	"github.com/pradanadp/simple-resto-app/features/menu"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.New()

	NullStockReport = menu.StockReportEntity{}
)

type menuService struct {
	repository menu.Repository
}

// GetStockReport implements menu.MenuService.
func (ms *menuService) GetStockReport(itemID string) (menu.StockReportEntity, error) {
	stockReport, err := ms.repository.GetStockReport(itemID)
	if err != nil {
		log.Error(err.Error())
		return NullStockReport, err
	}

	return stockReport, nil
}

func New(m menu.Repository) menu.Service {
	return &menuService{
		repository: m,
	}
}

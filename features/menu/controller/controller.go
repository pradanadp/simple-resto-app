package controller

import (
	"net/http"
	"strings"

	echo "github.com/labstack/echo/v4"
	"github.com/pradanadp/simple-resto-app/features/menu"
	"github.com/pradanadp/simple-resto-app/utils"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.New()
)

const (
	errItemIDRequired = "item ID is required"
	errItemNotFound   = "item not found"
)

type menuController struct {
	service menu.Service
}

// GetStockReport implements menu.MenuController.
func (mc *menuController) GetStockReport() echo.HandlerFunc {
	return func(c echo.Context) error {
		itemID := c.Param("item_id")
		if itemID == "" {
			log.Error(errItemIDRequired)

			return utils.Response(c, utils.RequestResponse{
				Code:  http.StatusBadRequest,
				Error: errItemIDRequired,
			})
		}

		stockReport, err := mc.service.GetStockReport(itemID)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				log.Error(errItemNotFound)

				return utils.Response(c, utils.RequestResponse{
					Code:  http.StatusNotFound,
					Error: errItemNotFound,
				})
			}
		}

		return utils.Response(c, utils.RequestResponse{
			Code: http.StatusOK,
			Data: stockReport,
		})
	}
}

func New(s menu.Service) menu.Controller {
	return &menuController{
		service: s,
	}
}

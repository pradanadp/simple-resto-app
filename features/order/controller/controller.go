package controller

import (
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"
	"github.com/pradanadp/simple-resto-app/app/middlewares"
	"github.com/pradanadp/simple-resto-app/features/order"
	"github.com/pradanadp/simple-resto-app/utils"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.New()
)

type orderController struct {
	service order.Service
}

// AddItemToCart implements order.Controller.
func (oc *orderController) AddItemToCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		customerID, err := middlewares.ExtractToken(c)
		if err != nil {
			log.Error("missing or malformed JWT")

			return utils.Response(c, utils.RequestResponse{
				Code:  http.StatusUnauthorized,
				Error: "missing or malformed JWT",
			})
		}

		itemID := c.QueryParam("item_id")
		quantity := c.QueryParam("quantity")
		qty, err := strconv.Atoi(quantity)
		if err != nil {
			log.Error("invalid quantity")

			return utils.Response(c, utils.RequestResponse{
				Code:  http.StatusBadRequest,
				Error: "invalid quantity",
			})
		}

		err = oc.service.AddItemToCart(itemID, customerID, uint(qty))
		if err != nil {
			log.Errorf("failed to add item to cart: %v", err)

			return utils.Response(c, utils.RequestResponse{
				Code:  http.StatusInternalServerError,
				Error: "failed to add item to cart",
			})
		}

		return utils.Response(c, utils.RequestResponse{
			Code:    http.StatusOK,
			Message: "Item added to cart successfully",
		})
	}
}

// GetMonthlyIncomeReport implements order.Controller.
func (oc *orderController) GetMonthlyIncomeReport() echo.HandlerFunc {
	return func(c echo.Context) error {
		yearMonth := c.QueryParam("year_month")

		if yearMonth == "" {
			return utils.Response(c, utils.RequestResponse{
				Code:  http.StatusBadRequest,
				Error: "year_month is required",
			})
		}

		incomeReport, err := oc.service.GetMonthlyIncomeReport(yearMonth)
		if err != nil {
			log.Errorf("failed to get monthly income report: %v", err)
			return utils.Response(c, utils.RequestResponse{
				Code:  http.StatusInternalServerError,
				Error: "failed to get monthly income report",
			})
		}

		return utils.Response(c, utils.RequestResponse{
			Code: http.StatusOK,
			Data: incomeReport,
		})
	}
}

// GetReceipt implements order.Controller.
func (oc *orderController) GetReceipt() echo.HandlerFunc {
	return func(c echo.Context) error {
		orderID := c.Param("order_id")

		if orderID == "" {
			return utils.Response(c, utils.RequestResponse{
				Code:  http.StatusBadRequest,
				Error: "order ID is required",
			})
		}

		receipt, err := oc.service.GetReceipt(orderID)
		if err != nil {
			log.Errorf("failed to get purchase receipt: %v", err)
			return utils.Response(c, utils.RequestResponse{
				Code:  http.StatusInternalServerError,
				Error: "failed to get purchase receipt",
			})
		}

		return utils.Response(c, utils.RequestResponse{
			Code: http.StatusOK,
			Data: receipt,
		})
	}
}

// GetWeeklyIncomeReport implements order.Controller.
func (oc *orderController) GetWeeklyIncomeReport() echo.HandlerFunc {
	return func(c echo.Context) error {
		startDate := c.QueryParam("start_date")
		endDate := c.QueryParam("end_date")

		if startDate == "" || endDate == "" {
			return utils.Response(c, utils.RequestResponse{
				Code:  http.StatusBadRequest,
				Error: "start_date and end_date are required",
			})
		}

		incomeReport, err := oc.service.GetWeeklyIncomeReport(startDate, endDate)
		if err != nil {
			log.Errorf("failed to get weekly income report: %v", err)
			return utils.Response(c, utils.RequestResponse{
				Code:  http.StatusInternalServerError,
				Error: "failed to get weekly income report",
			})
		}

		return utils.Response(c, utils.RequestResponse{
			Code: http.StatusOK,
			Data: incomeReport,
		})
	}
}

// RemoveItemFromCart implements order.Controller.
func (oc *orderController) RemoveItemFromCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		itemID := c.Param("item_id")

		if itemID == "" {
			return utils.Response(c, utils.RequestResponse{
				Code:  http.StatusBadRequest,
				Error: "item ID is required",
			})
		}

		err := oc.service.RemoveItemFromCart(itemID)
		if err != nil {
			log.Errorf("failed to remove item from cart: %v", err)
			return utils.Response(c, utils.RequestResponse{
				Code:  http.StatusInternalServerError,
				Error: "failed to remove item from cart",
			})
		}

		return utils.Response(c, utils.RequestResponse{
			Code:    http.StatusOK,
			Message: "Item removed from cart successfully",
		})
	}
}

func New(s order.Service) order.Controller {
	return &orderController{
		service: s,
	}
}

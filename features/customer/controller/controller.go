package controller

import (
	"net/http"
	"strings"

	echo "github.com/labstack/echo/v4"
	"github.com/pradanadp/simple-resto-app/features/customer"
	"github.com/pradanadp/simple-resto-app/utils"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.New()
)

type customerController struct {
	service customer.Service
}

// Login implements customer.Controller.
func (cc *customerController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		request := LoginRequest{}
		err := c.Bind(&request)
		if err != nil {
			log.Error("controller - error on bind request")
			return utils.Response(c, utils.RequestResponse{
				Code:  http.StatusBadRequest,
				Error: "Bad request: unable to parse request body",
			})
		}

		resp, token, err := cc.service.Login(RequestToEntity(request))
		if err != nil {
			switch {
			case strings.Contains(err.Error(), "invalid email format"):
				log.Error("bad request, invalid email format")
				return utils.Response(c, utils.RequestResponse{
					Code:  http.StatusBadRequest,
					Error: "Bad request: invalid email format",
				})
			case strings.Contains(err.Error(), "password cannot be empty"):
				log.Error("bad request, password cannot be empty")
				return utils.Response(c, utils.RequestResponse{
					Code:  http.StatusBadRequest,
					Error: "Bad request: password cannot be empty",
				})
			case strings.Contains(err.Error(), "invalid email and password"):
				log.Error("bad request, invalid email and password")
				return utils.Response(c, utils.RequestResponse{
					Code:  http.StatusBadRequest,
					Error: "Bad request: invalid email and password",
				})
			case strings.Contains(err.Error(), "password does not match"):
				log.Error("bad request, password does not match")
				return utils.Response(c, utils.RequestResponse{
					Code:  http.StatusBadRequest,
					Error: "Bad request: password does not match",
				})
			case strings.Contains(err.Error(), "no row affected"):
				log.Error("customer not found")
				return utils.Response(c, utils.RequestResponse{
					Code:  http.StatusNotFound,
					Error: "Customer not found",
				})
			case strings.Contains(err.Error(), "error while creating jwt token"):
				log.Error("internal server error, error while creating jwt token")
				return utils.Response(c, utils.RequestResponse{
					Code:  http.StatusInternalServerError,
					Error: "Internal server error",
				})
			default:
				log.Error("internal server error")
				return utils.Response(c, utils.RequestResponse{
					Code:  http.StatusInternalServerError,
					Error: "Internal server error",
				})
			}
		}

		// Respond with the successful login response and token
		return utils.Response(c, utils.RequestResponse{
			Code: http.StatusOK,
			Data: map[string]interface{}{
				"response": resp,
				"token":    token,
			},
		})
	}
}

// Register implements customer.Controller.
func (cc *customerController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := RegisterRequest{}
		err := c.Bind(&req)
		if err != nil {
			log.Error("controller - error on bind request")
			return utils.Response(c, utils.RequestResponse{
				Code:  http.StatusBadRequest,
				Error: "Bad request: invalid request payload",
			})
		}

		customerEntity := RequestToEntity(req)
		customer, err := cc.service.Register(customerEntity)
		if err != nil {
			log.Errorf("failed to register new customer: %v", err)
			return utils.Response(c, utils.RequestResponse{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
		}

		// Respond with the successful registration response
		return utils.Response(c, utils.RequestResponse{
			Code:    http.StatusCreated,
			Data:    customer,
			Message: "Register Successful",
		})
	}
}

func New(s customer.Service) customer.Controller {
	return &customerController{
		service: s,
	}
}

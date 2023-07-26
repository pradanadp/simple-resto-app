package utils

import "github.com/labstack/echo/v4"

type RequestResponse struct {
	Code    int         `json:"code,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
}

func Response(c echo.Context, r RequestResponse) error {
	return c.JSON(r.Code, r)
}

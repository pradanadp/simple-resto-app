package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pradanadp/simple-resto-app/app/middlewares"
	cc "github.com/pradanadp/simple-resto-app/features/customer/controller"
	cd "github.com/pradanadp/simple-resto-app/features/customer/repository"
	cs "github.com/pradanadp/simple-resto-app/features/customer/service"
	mc "github.com/pradanadp/simple-resto-app/features/menu/controller"
	md "github.com/pradanadp/simple-resto-app/features/menu/repository"
	ms "github.com/pradanadp/simple-resto-app/features/menu/service"
	oc "github.com/pradanadp/simple-resto-app/features/order/controller"
	od "github.com/pradanadp/simple-resto-app/features/order/repository"
	os "github.com/pradanadp/simple-resto-app/features/order/service"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	initCustomerRouter(db, e)
	initMenuRouter(db, e)
}

func initCustomerRouter(db *gorm.DB, e *echo.Echo) {
	repo := cd.New(db)
	validate := validator.New()
	service := cs.New(repo, validate)
	handler := cc.New(service)

	e.POST("/register", handler.Register())
	e.POST("/login", handler.Login())
}

func initMenuRouter(db *gorm.DB, e *echo.Echo) {
	repo := md.New(db)
	service := ms.New(repo)
	handler := mc.New(service)

	e.GET("/report/stock", handler.GetStockReport())
}

func initOrderRouter(db *gorm.DB, e *echo.Echo) {
	repo := od.New(db)
	service := os.New(repo)
	handler := oc.New(service)

	e.POST("/cart", handler.AddItemToCart(), middlewares.JWTMiddleware())
	e.DELETE("/cart", handler.RemoveItemFromCart(), middlewares.JWTMiddleware())
	e.GET("/report/income/weekly", handler.GetWeeklyIncomeReport())
	e.GET("/report/income/monthly", handler.GetMonthlyIncomeReport())
	e.GET("/receipt", handler.GetReceipt())
}

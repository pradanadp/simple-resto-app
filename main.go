package main

import (
	"github.com/labstack/echo/v4"
	"github.com/pradanadp/simple-resto-app/app/config"
	"github.com/pradanadp/simple-resto-app/app/database"
	"github.com/pradanadp/simple-resto-app/app/router"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := database.InitDB(cfg)
	router.InitRouter(db, e)
	e.Logger.Fatal(e.Start(":8080"))
}

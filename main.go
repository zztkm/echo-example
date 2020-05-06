package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/zztkm/echo-example/handler"
)

func main() {
	// create echo instance
	e := echo.New()

	// setting middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// routing
	e.GET("/hello", handler.MainPage())

	// start server
	e.Logger.Fatal(e.Start(":1323"))
}

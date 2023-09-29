package main

import (
	"prakerja/config"
	"prakerja/controller"

	"github.com/labstack/echo/v4"
)

func main() {
	// Create HTTP server
	e := echo.New()

	// Connect To Database
	config.DatabaseInit()
	gorm := config.DB()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	dbGorm.Ping()

	bookRoute := e.Group("/book")
	bookRoute.POST("/", controller.CreateBook)
	bookRoute.GET("/:id", controller.GetBook)
	bookRoute.PUT("/:id", controller.UpdateBook)
	bookRoute.DELETE("/:id", controller.DeleteBook)

	e.Logger.Fatal(e.Start(":8082"))
}

package main

import (
	"Test-Task/pkg"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	server := echo.New()
	server.Use(pkg.Middleware)
	server.GET("/Days", pkg.Data)
	err := server.Start(":3000")
	if err != nil {
		log.Fatal(err)
	}
}

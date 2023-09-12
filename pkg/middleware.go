package pkg

import (
	"github.com/labstack/echo/v4"
	"log"
)

func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role := c.Request().Header.Get("User-Role")
		if role == "admin" {
			log.Println("red button users detecteed")
		}
		return next(c)
	}
}

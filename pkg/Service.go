package pkg

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func Data(c echo.Context) error {
	Time := time.Now()
	date := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	diff := date.Sub(Time)
	return c.String(http.StatusOK, fmt.Sprintf("diff is: %f days", int(diff.Hours()/24)))
}

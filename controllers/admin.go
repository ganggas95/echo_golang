package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

type Admin struct {
}

func (a *Admin) Index() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "admin_index", nil)
	}
}

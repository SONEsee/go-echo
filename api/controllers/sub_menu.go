package controllers

import (
	"net/http"

	"github.com/SONEsee/go-echo/api/presenters"
	"github.com/SONEsee/go-echo/api/services"
	"github.com/labstack/echo/v4"
)

func GetSubllMenu(c echo.Context) error {
	SubMenu, err := services.GateAllWhitSubmenu(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess(SubMenu))
}

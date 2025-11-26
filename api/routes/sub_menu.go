package routes

import (
	"github.com/SONEsee/go-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutesSubmenu(routes *echo.Group) {
	routes.GET("/suball-menu/", controllers.GetSubllMenu)
}

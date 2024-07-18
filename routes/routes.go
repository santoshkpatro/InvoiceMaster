package routes

import (
	"InvoiceMaster/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.POST("/api/user/register", controllers.RegisterUser)
}

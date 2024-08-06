package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"icon-pln/internal/delivery/http/handler"
)

func NewRouter(userHandler *handler.DashboardHandler) *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/api/dashboard", userHandler.GetReportDashboard)

	return e
}

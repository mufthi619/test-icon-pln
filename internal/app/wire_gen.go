// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/labstack/echo/v4"
	"icon-pln/internal/config"
	"icon-pln/internal/delivery/http/handler"
	"icon-pln/internal/delivery/router"
	"icon-pln/internal/repository/http"
	"icon-pln/internal/usecase"
)

// Injectors from wire.go:

func BootstrapApp(cfg *config.Config) (*echo.Echo, func(), error) {
	dashboardRepository := http.NewDashboardRepository(cfg)
	dashboardUseCase := usecase.NewDashboardUseCase(dashboardRepository)
	dashboardHandler := handler.NewDashboardHandler(dashboardUseCase)
	echoEcho := router.NewRouter(dashboardHandler)
	return echoEcho, func() {
	}, nil
}

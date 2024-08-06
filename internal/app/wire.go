//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"icon-pln/internal/config"
	"icon-pln/internal/delivery/http/handler"
	"icon-pln/internal/delivery/router"
	"icon-pln/internal/repository/http"
	"icon-pln/internal/usecase"
)

func BootstrapApp(cfg *config.Config) (*echo.Echo, func(), error) {
	wire.Build(
		http.NewDashboardRepository,
		usecase.NewDashboardUseCase,
		handler.NewDashboardHandler,
		router.NewRouter,
	)
	return nil, nil, nil
}

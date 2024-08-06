package handler

import (
	"github.com/labstack/echo/v4"
	"icon-pln/internal/domain/dashboard"
	"icon-pln/internal/usecase"
	"net/http"
)

type DashboardHandler struct {
	dashboardUseCase usecase.DashboardUseCase
}

func NewDashboardHandler(dashboardUseCase usecase.DashboardUseCase) *DashboardHandler {
	return &DashboardHandler{
		dashboardUseCase: dashboardUseCase,
	}
}

func (h *DashboardHandler) GetReportDashboard(ctx echo.Context) error {
	var filter dashboard.GetDashboardFilter
	if err := ctx.Bind(&filter); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Failed ! Invalid body",
			"data":    nil,
		})
	}

	resp, msg, err := h.dashboardUseCase.GetReportDashboard(filter)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": msg,
			"data":    resp,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": msg,
		"data":    resp,
	})
}

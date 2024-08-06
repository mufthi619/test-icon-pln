package usecase

import "icon-pln/internal/domain/dashboard"

type DashboardUseCase interface {
	GetReportDashboard(filter dashboard.GetDashboardFilter) ([]dashboard.Response, string, error)
}

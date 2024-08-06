package repository

import "icon-pln/internal/domain/dashboard"

type DashboardRepository interface {
	GetBookingList(filter dashboard.GetDashboardFilter) (finalResponse dashboard.BookingListAPI, err error)
	GetConsumptionType(filter dashboard.GetDashboardFilter) (finalResponse dashboard.ConsumptionTypeAPIs, err error)
}

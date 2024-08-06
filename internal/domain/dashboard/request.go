package dashboard

import "time"

type DateRangeFilter int
type Environment int

const (
	Today DateRangeFilter = iota + 1
	Yesterday
	LastWeek
	CurrentWeek
	LastMonth
	CurrentMonth
	LastYear
	CurrentYear
	ChoseDate
	All
)

type GetDashboardFilter struct {
	RangeFilter  DateRangeFilter `json:"range_filter" query:"range_filter"`
	StartDateStr *string         `json:"start_date" query:"start_date"`
	EndDateStr   *string         `json:"end_date" query:"end_date"`
	StartDate    *time.Time      `json:"-"`
	EndDate      *time.Time      `json:"-"`
}

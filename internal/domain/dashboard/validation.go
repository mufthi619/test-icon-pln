package dashboard

import (
	"fmt"
	"time"
)

func (data *GetDashboardFilter) ValidateAndParse() (string, error) {

	if data.RangeFilter < Today || data.RangeFilter > ChoseDate {
		//return fmt.Sprintf("invalid range filter: must be between %d and %d", Today, ChoseDate), fmt.Errorf("invalid range filter: must be between %d and %d", Today, ChoseDate)
		data.RangeFilter = CurrentMonth
	}

	switch data.RangeFilter {
	case ChoseDate:
		if data.StartDateStr == nil || data.EndDateStr == nil {
			return fmt.Sprintf("start_date and end_date are required when range_filter is ChoseDate"), fmt.Errorf("start_date and end_date are required when range_filter is ChoseDate")
		}

		startDate, err := time.Parse("2006-01-02", *data.StartDateStr)
		if err != nil {
			return fmt.Sprintf("invalid start_date format: %v", err), fmt.Errorf("invalid start_date format: %v", err)
		}
		data.StartDate = &startDate

		endDate, err := time.Parse("2006-01-02", *data.EndDateStr)
		if err != nil {
			return "", fmt.Errorf("invalid end_date format: %v", err)
		}
		data.EndDate = &endDate

		if startDate.After(endDate) {
			return "", fmt.Errorf("start_date must be before or equal to end_date")
		}
		break
	case Today:
		now := time.Now()
		year, month, day := now.Date()
		startAt := time.Date(year, month, day, 0, 0, 0, 0, now.Location())
		endAt := startAt.AddDate(0, 0, 1).Add(-1 * time.Nanosecond)

		data.StartDate = &startAt
		data.EndDate = &endAt
		break
	case Yesterday:
		now := time.Now()
		yesterday := now.AddDate(0, 0, -1)

		year, month, day := yesterday.Date()
		startAt := time.Date(year, month, day, 0, 0, 0, 0, now.Location())
		endAt := startAt.AddDate(0, 0, 1).Add(-1 * time.Nanosecond)

		data.StartDate = &startAt
		data.EndDate = &endAt
		break
	case CurrentWeek:
		now := time.Now()
		weekday := now.Weekday()

		if weekday == time.Sunday {
			weekday = 7
		} else {
			weekday--
		}

		startAt := now.AddDate(0, 0, -int(weekday)).Truncate(24 * time.Hour)
		endAt := startAt.AddDate(0, 0, 7).Add(-1 * time.Nanosecond)

		data.StartDate = &startAt
		data.EndDate = &endAt
		break
	case LastWeek:
		now := time.Now()
		weekday := now.Weekday()

		if weekday == time.Sunday {
			weekday = 7
		} else {
			weekday--
		}

		startAt := now.AddDate(0, 0, -int(weekday)-7).Truncate(24 * time.Hour)
		endAt := startAt.AddDate(0, 0, 7).Add(-1 * time.Nanosecond)

		data.StartDate = &startAt
		data.EndDate = &endAt
		break
	case CurrentMonth:
		now := time.Now()
		currentYear, currentMonth, _ := now.Date()
		startAt := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, now.Location())
		endAt := startAt.AddDate(0, 1, 0).Add(-1 * time.Nanosecond)

		data.StartDate = &startAt
		data.EndDate = &endAt
		break
	case LastMonth:
		now := time.Now()
		currentYear, currentMonth, _ := now.Date()
		lastMonth := currentMonth - 1
		lastMonthYear := currentYear

		if currentMonth == time.January {
			lastMonth = time.December
			lastMonthYear--
		}

		startAt := time.Date(lastMonthYear, lastMonth, 1, 0, 0, 0, 0, now.Location())
		endAt := startAt.AddDate(0, 1, 0).Add(-1 * time.Nanosecond)

		data.StartDate = &startAt
		data.EndDate = &endAt
		break
	case CurrentYear:
		now := time.Now()
		currentYear := now.Year()
		startAt := time.Date(currentYear, time.January, 1, 0, 0, 0, 0, now.Location())
		endAt := time.Date(currentYear, time.December, 31, 23, 59, 59, 999999999, now.Location())

		data.StartDate = &startAt
		data.EndDate = &endAt
		break
	case LastYear:
		now := time.Now()
		lastYear := now.Year() - 1
		startAt := time.Date(lastYear, time.January, 1, 0, 0, 0, 0, now.Location())
		endAt := time.Date(lastYear, time.December, 31, 23, 59, 59, 999999999, now.Location())

		data.StartDate = &startAt
		data.EndDate = &endAt
		break
	}

	return "Successfully", nil
}

func (data *GetDashboardFilter) SetFilterDateToNil() {
	data.StartDate = nil
	data.EndDate = nil
}

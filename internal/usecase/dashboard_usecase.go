package usecase

import (
	"github.com/labstack/gommon/log"
	"icon-pln/internal/domain/dashboard"
	"icon-pln/internal/repository"
	"icon-pln/internal/transformer"
	"sync"
)

type dashboardUseCase struct {
	dashboardRepo repository.DashboardRepository
}

func NewDashboardUseCase(dashboardRepo repository.DashboardRepository) DashboardUseCase {
	return &dashboardUseCase{
		dashboardRepo: dashboardRepo,
	}
}

func (s *dashboardUseCase) GetReportDashboard(filter dashboard.GetDashboardFilter) ([]dashboard.Response, string, error) {
	var finalResponse []dashboard.Response
	const internalProblem = "Failed ! There's some trouble in our system, please try again"

	//Validate Data
	msg, err := filter.ValidateAndParse()
	if err != nil {
		log.Info("[GetReportDashboard][Flag-1] | Failed on GetReportDashboard, err -> ", err)
		return finalResponse, msg, err
	}

	//Retrieve Data From Repository
	consumptionTypeChan := make(chan struct {
		data dashboard.ConsumptionTypeAPIs
		err  error
	})
	bookingListChan := make(chan struct {
		data dashboard.BookingListAPI
		err  error
	})
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		resp, err := s.dashboardRepo.GetBookingList(filter)
		if err != nil {
			bookingListChan <- struct {
				data dashboard.BookingListAPI
				err  error
			}{data: resp, err: err}
			return
		}
		bookingListChan <- struct {
			data dashboard.BookingListAPI
			err  error
		}{data: resp, err: nil}
		return
	}()
	go func() {
		defer wg.Done()
		resp, err := s.dashboardRepo.GetConsumptionType(filter)
		if err != nil {
			consumptionTypeChan <- struct {
				data dashboard.ConsumptionTypeAPIs
				err  error
			}{data: resp, err: err}
		}
		consumptionTypeChan <- struct {
			data dashboard.ConsumptionTypeAPIs
			err  error
		}{data: resp, err: nil}
		return
	}()
	go func() {
		wg.Wait()
		close(consumptionTypeChan)
		close(bookingListChan)
	}()

	consumptionType := <-consumptionTypeChan
	bookingList := <-bookingListChan

	if consumptionType.err != nil {
		log.Info("[GetReportDashboard][Flag-2] | Failed on consumptionType, err -> ", consumptionType.err)
		return finalResponse, internalProblem, err
	}
	if bookingList.err != nil {
		log.Info("[GetReportDashboard][Flag-3] | Failed on bookingList, err -> ", bookingList.err)
		return finalResponse, internalProblem, err
	}

	return transformer.TransformBookingListToResponse(consumptionType.data, bookingList.data), "Successfully", nil
}

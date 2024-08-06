package http

import (
	"encoding/json"
	"errors"
	"icon-pln/internal/config"
	"icon-pln/internal/domain/dashboard"
	"icon-pln/internal/repository"
	"net/http"
)

type dashboardRepository struct {
	client *http.Client
}

func NewDashboardRepository(cfg *config.Config) repository.DashboardRepository {
	return &dashboardRepository{
		client: &http.Client{
			Timeout: cfg.ExternalAPI.Timeout,
		},
	}
}

func (r *dashboardRepository) GetBookingList(filter dashboard.GetDashboardFilter) (finalResponse dashboard.BookingListAPI, err error) {
	url := "https://66876cc30bc7155dc017a662.mockapi.io/api/dummy-data/bookingList"

	resp, err := r.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = errors.New("failed ! HTTP Response not 200")
		return
	}

	if err := json.NewDecoder(resp.Body).Decode(&finalResponse); err != nil {
		return nil, err
	}

	return
}

func (r *dashboardRepository) GetConsumptionType(filter dashboard.GetDashboardFilter) (finalResponse dashboard.ConsumptionTypeAPIs, err error) {
	url := "https://6686cb5583c983911b03a7f3.mockapi.io/api/dummy-data/masterJenisKonsumsi"

	resp, err := r.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = errors.New("failed ! HTTP Response not 200")
	}

	if err := json.NewDecoder(resp.Body).Decode(&finalResponse); err != nil {
		return nil, err
	}

	return
}

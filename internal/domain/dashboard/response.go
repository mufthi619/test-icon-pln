package dashboard

import "time"

type (
	ConsumptionTypeAPI struct {
		CreatedAt time.Time `json:"createdAt"`
		Name      string    `json:"name"`
		MaxPrice  int       `json:"maxPrice"`
		ID        string    `json:"id"`
	}

	ConsumptionTypeAPIs []ConsumptionTypeAPI
)

type (
	BookingListAPI []BookingAPI

	BookingAPI struct {
		BookingDate     time.Time         `json:"bookingDate"`
		OfficeName      string            `json:"officeName"`
		StartTime       time.Time         `json:"startTime"`
		EndTime         time.Time         `json:"endTime"`
		ListConsumption []ListConsumption `json:"listConsumption"`
		Participants    int               `json:"participants"`
		RoomName        string            `json:"roomName"`
		ID              string            `json:"id"`
	}

	ListConsumption struct {
		Name string `json:"name"`
	}
)

type (
	Response struct {
		OfficeName   string         `json:"office_name"`
		RoomResponse []RoomResponse `json:"room_response"`
	}

	RoomResponse struct {
		RoomName           string                    `json:"room_name"`
		UsagePercentage    float64                   `json:"usage_percentage"`
		ConsumptionNominal float64                   `json:"consumption_nominal"`
		ConsumptionType    []ConsumptionTypeResponse `json:"consumption_type"`
	}

	ConsumptionTypeResponse struct {
		Name  string `json:"name"`
		Total int    `json:"total"`
	}
)

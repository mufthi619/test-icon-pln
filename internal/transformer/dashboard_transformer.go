package transformer

import "icon-pln/internal/domain/dashboard"

func TransformBookingListToResponse(consumptionType dashboard.ConsumptionTypeAPIs, bookingList dashboard.BookingListAPI) (finalResponse []dashboard.Response) {
	officeRoomMap := make(map[string]map[string]bool)
	for _, booking := range bookingList {
		if _, ok := officeRoomMap[booking.OfficeName]; !ok {
			officeRoomMap[booking.OfficeName] = make(map[string]bool)
		}
		officeRoomMap[booking.OfficeName][booking.RoomName] = true
	}

	for officeName, rooms := range officeRoomMap {
		roomResponses := make([]dashboard.RoomResponse, 0, len(rooms))
		for roomName := range rooms {
			roomResponses = append(roomResponses, dashboard.RoomResponse{
				RoomName:        roomName,
				ConsumptionType: []dashboard.ConsumptionTypeResponse{},
			})
		}
		finalResponse = append(finalResponse, dashboard.Response{
			OfficeName:   officeName,
			RoomResponse: roomResponses,
		})
	}

	consumptionPrices := make(map[string]int)
	for _, ct := range consumptionType {
		consumptionPrices[ct.Name] = ct.MaxPrice
	}

	for i, office := range finalResponse {
		for j, room := range office.RoomResponse {
			totalHours := 0.0
			totalConsumption := 0.0
			consumptionCounts := make(map[string]int)

			for _, booking := range bookingList {
				if booking.OfficeName == office.OfficeName && booking.RoomName == room.RoomName {
					duration := booking.EndTime.Sub(booking.StartTime).Hours()
					totalHours += duration

					for _, consumption := range booking.ListConsumption {
						if price, ok := consumptionPrices[consumption.Name]; ok {
							totalConsumption += float64(price * booking.Participants)
							consumptionCounts[consumption.Name] += booking.Participants
						}
					}
				}
			}

			// Calculate usage percentage (assuming 8 hours per day and 30 days in a month)
			usagePercentage := (totalHours / (8 * 30)) * 100

			finalResponse[i].RoomResponse[j].UsagePercentage = usagePercentage
			finalResponse[i].RoomResponse[j].ConsumptionNominal = totalConsumption

			// Add individual consumption type responses
			for name, count := range consumptionCounts {
				finalResponse[i].RoomResponse[j].ConsumptionType = append(
					finalResponse[i].RoomResponse[j].ConsumptionType,
					dashboard.ConsumptionTypeResponse{
						Name:  name,
						Total: count,
					},
				)
			}
		}
	}

	return
}

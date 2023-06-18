package startup

import (
	"Booking/flight-service/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var flights = []*domain.Flight{
	{
		ID:                primitive.NewObjectID(),
		Departure:         "20-06-2023",
		DeparturePlace:    "Beograd",
		ArrivalPlace:      "Bec",
		Price:             36,
		NumberOfFreeSeats: 25,
	},
	{
		ID:                primitive.NewObjectID(),
		Departure:         "20-06-2023",
		DeparturePlace:    "Beograd",
		ArrivalPlace:      "Rejkjavik",
		Price:             55,
		NumberOfFreeSeats: 15,
	},
	{
		ID:                primitive.NewObjectID(),
		Departure:         "21-06-2023",
		DeparturePlace:    "London",
		ArrivalPlace:      "Beograd",
		Price:             70,
		NumberOfFreeSeats: 5,
	},
}

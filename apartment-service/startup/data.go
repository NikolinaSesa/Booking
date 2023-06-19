package startup

import (
	"github.com/NikolinaSesa/Booking/apartment-service/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var apartments = []*domain.Apartment{
	{
		Id:                   primitive.NewObjectID(),
		HostId:               primitive.NewObjectID(),
		Name:                 "Lux apartmani",
		Location:             "Zlatibor",
		Benefits:             "WiFi, Parking",
		MinGuestsNumber:      2,
		MaxGuestsNumber:      6,
		AutomaticReservation: false,
		GeneralPrice:         "300",
		Ratings: []domain.Rating{
			{
				Rating: 1,
			},
		},
	},
	{
		Id:                   primitive.NewObjectID(),
		HostId:               primitive.NewObjectID(),
		Name:                 "Apartmani Sunce",
		Location:             "Zlatibor",
		Benefits:             "WiFi, Parking, Bazen",
		MinGuestsNumber:      2,
		MaxGuestsNumber:      6,
		AutomaticReservation: false,
		GeneralPrice:         "250",
		Ratings: []domain.Rating{
			{
				Rating: 3,
			},
		},
	},
}

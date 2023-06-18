package startup

import (
	"github.com/NikolinaSesa/Booking/apartment-service/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var apartments = []*domain.Apartment{
	{
		Id:                   getObjectId("648f2fb76a8458e27e56284e"),
		HostId:               getObjectId("648f314999863f768f378307"),
		Name:                 "Lux apartmani",
		Location:             "Zlatibor",
		Benefits:             "WiFi, Parking",
		MinGuestsNumber:      2,
		MaxGuestsNumber:      6,
		AutomaticReservation: false,
	},
	{
		Id:                   getObjectId("648f2fb76a8458e27e562850"),
		HostId:               getObjectId("648f314999863f768f378307"),
		Name:                 "Apartmani Sunce",
		Location:             "Zlatibor",
		Benefits:             "WiFi, Parking, Bazen",
		MinGuestsNumber:      2,
		MaxGuestsNumber:      6,
		AutomaticReservation: false,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}

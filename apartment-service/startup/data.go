package startup

import (
	"github.com/NikolinaSesa/Booking/apartment-service/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var apartments = []*domain.Apartment{
	{
		Id:                   primitive.NewObjectID(),
		HostId:               getObjectId("648f314999863f768f378304"),
		Name:                 "Lux apartmani",
		Location:             "Zlatibor",
		Benefits:             "WiFi",
		MinGuestsNumber:      2,
		MaxGuestsNumber:      6,
		AutomaticReservation: false,
		GeneralPrice:         150,
	},
	{
		Id:                   primitive.NewObjectID(),
		HostId:               primitive.NewObjectID(),
		Name:                 "Apartmani Sunce",
		Location:             "Zlatibor",
		Benefits:             "WiFi,Parking,Bazen",
		MinGuestsNumber:      2,
		MaxGuestsNumber:      6,
		AutomaticReservation: false,
		GeneralPrice:         200,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}

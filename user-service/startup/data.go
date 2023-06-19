package startup

import (
	"github.com/NikolinaSesa/Booking/user-service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var users = []*domain.User{
	{
		Id:        getObjectId("648f314999863f768f378304"),
		FirstName: "Nikolina",
		LastName:  "Sesa",
		Email:     "nikolinasesaa23@gmail.com",
		Password:  "123",
		Username:  "nikolina",
		Address:   "Nikole Tesle 45, Novi Sad",
		Role:      "HOST",
		Mark:      "true",
		Ratings: []domain.Rating{
			{
				Rating: 1,
			},
		},
		AvgRating: 1.0,
	},
	{
		Id:        getObjectId("648f314999863f768f378305"),
		FirstName: "Nenad",
		LastName:  "Joldic",
		Email:     "nenad123@gmail.com",
		Password:  "123",
		Username:  "nenad",
		Address:   "Micurinova 3, Novi Sad",
		Role:      "GUEST",
		Mark:      "false",
	},
	{
		Id:        primitive.NewObjectID(),
		FirstName: "Bogdan",
		LastName:  "Blagojevic",
		Email:     "boki123@gmail.com",
		Password:  "123",
		Username:  "boki",
		Address:   "Rumenacka 55, Novi Sad",
		Role:      "GUEST",
		Mark:      "false",
	},
	{
		Id:        primitive.NewObjectID(),
		FirstName: "Srdjan",
		LastName:  "Tosic",
		Email:     "srkitosic123@gmail.com",
		Password:  "123",
		Username:  "srki",
		Address:   "Bele Njive 66, Novi Sad",
		Role:      "HOST",
		Mark:      "false",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}

var reservations = []*domain.Reservation{
	{
		ID:           primitive.NewObjectID(),
		GuestID:      getObjectId("648f314999863f768f378305"),
		ApartmentID:  primitive.NewObjectID(),
		StartDate:    "30-5-2023",
		EndDate:      "24-6-2023",
		GuestsNumber: 50,
	},
	{
		ID:           primitive.NewObjectID(),
		GuestID:      primitive.NewObjectID(),
		ApartmentID:  primitive.NewObjectID(),
		StartDate:    "30-5-2023",
		EndDate:      "24-6-2023",
		GuestsNumber: 50,
	},
	{
		ID:           primitive.NewObjectID(),
		GuestID:      getObjectId("648f314999863f768f378305"),
		ApartmentID:  primitive.NewObjectID(),
		StartDate:    "30-5-2023",
		EndDate:      "24-6-2023",
		GuestsNumber: 50,
	},
	{
		ID:           primitive.NewObjectID(),
		GuestID:      primitive.NewObjectID(),
		ApartmentID:  primitive.NewObjectID(),
		StartDate:    "30-5-2023",
		EndDate:      "24-6-2023",
		GuestsNumber: 50,
	},
	{
		ID:           primitive.NewObjectID(),
		GuestID:      getObjectId("648f314999863f768f378305"),
		ApartmentID:  primitive.NewObjectID(),
		StartDate:    "30-5-2023",
		EndDate:      "24-6-2023",
		GuestsNumber: 50,
	},
	{
		ID:           primitive.NewObjectID(),
		GuestID:      getObjectId("648f314999863f768f378305"),
		ApartmentID:  primitive.NewObjectID(),
		StartDate:    "30-5-2023",
		EndDate:      "24-6-2023",
		GuestsNumber: 50,
	},
	{
		ID:           primitive.NewObjectID(),
		GuestID:      getObjectId("648f314999863f768f378305"),
		ApartmentID:  primitive.NewObjectID(),
		StartDate:    "30-5-2023",
		EndDate:      "24-6-2023",
		GuestsNumber: 50,
	},
	{
		ID:           primitive.NewObjectID(),
		GuestID:      primitive.NewObjectID(),
		ApartmentID:  primitive.NewObjectID(),
		StartDate:    "30-5-2023",
		EndDate:      "24-6-2023",
		GuestsNumber: 50,
	},
}

package startup

import (
	"github.com/NikolinaSesa/Booking/user-service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var users = []*domain.User{
	{
		Id:        primitive.NewObjectID(),
		FirstName: "Nikolina",
		LastName:  "Sesa",
		Email:     "nikolinasesaa23@gmail.com",
		Password:  "123",
		Username:  "nikolina",
		Address:   "Nikole Tesle 45, Novi Sad",
		Role:      "GUEST",
	},
	{
		Id:        primitive.NewObjectID(),
		FirstName: "Nenad",
		LastName:  "Joldic",
		Email:     "nenad123@gmail.com",
		Password:  "123",
		Username:  "nenad",
		Address:   "Micurinova 3, Novi Sad",
		Role:      "GUEST",
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
	},
}

/*
func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
*/

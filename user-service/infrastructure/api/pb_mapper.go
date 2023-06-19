package api

import (
	"github.com/NikolinaSesa/Booking/user-service/domain"
	pb "github.com/NikolinaSesa/Booking/user-service/proto"
	"strconv"
)

func mapUser(user *domain.User) *pb.User {
	userPb := &pb.User{
		Id:        user.Id.Hex(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role,
		Mark:      user.Mark,
	}
	return userPb
}

func mapApartment(apartment *domain.Apartment) *pb.Apartment {
	apartmentPb := &pb.Apartment{
		Id:           apartment.Id.Hex(),
		Name:         apartment.Name,
		Location:     apartment.Location,
		Benefits:     apartment.Benefits,
		GeneralPrice: strconv.Itoa(apartment.GeneralPrice),
	}
	return apartmentPb
}

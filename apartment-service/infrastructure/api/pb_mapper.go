package api

import (
	"github.com/NikolinaSesa/Booking/apartment-service/domain"
	pb "github.com/NikolinaSesa/Booking/apartment-service/proto"
)

func mapApartment(apartment *domain.Apartment) *pb.Apartment {
	apartmentPb := &pb.Apartment{
		Id:       apartment.Id.Hex(),
		Name:     apartment.Name,
		Location: apartment.Location,
	}
	return apartmentPb
}

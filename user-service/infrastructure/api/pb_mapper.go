package api

import (
	"github.com/NikolinaSesa/Booking/user-service/domain"
	pb "github.com/NikolinaSesa/Booking/user-service/proto"
<<<<<<< HEAD
	"strconv"
=======
	"go.mongodb.org/mongo-driver/bson/primitive"
>>>>>>> origin/update-proto
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

<<<<<<< HEAD
func mapApartment(apartment *domain.Apartment) *pb.Apartment {
=======
func mapHostRating(user *domain.User) *pb.Host {
	hostRatingPb := &pb.Host{
		Id:        user.Id.Hex(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
	return hostRatingPb
}

func mapApartmentRating(apartment *domain.Apartment) *pb.Apartment {
>>>>>>> origin/update-proto
	apartmentPb := &pb.Apartment{
		Id:           apartment.Id.Hex(),
		Name:         apartment.Name,
		Location:     apartment.Location,
		Benefits:     apartment.Benefits,
<<<<<<< HEAD
		GeneralPrice: strconv.Itoa(apartment.GeneralPrice),
=======
		GeneralPrice: "test",
	}
	return apartmentPb
}

func mapHostRating2(hostRatingPb *pb.HostsRating) *domain.HostRating {
	objID, err := primitive.ObjectIDFromHex(hostRatingPb.HostId)
	if err != nil {
		panic(err)
	}
	hostRating := &domain.HostRating{
		HostId: objID,
		//UserId:        primitive.ObjectIDFromHex(hostRatingPb.UserId),
		UserFirstName: hostRatingPb.UserFirstName,
		UserLastName:  hostRatingPb.UserLastName,
		Rating:        hostRatingPb.Rating,
	}

	return hostRating
}

func mapApartmentRating2(apartmentRatingPb *pb.ApartmentsRating) *domain.ApartmentRating {
	objID, err := primitive.ObjectIDFromHex(apartmentRatingPb.ApartmentId)
	if err != nil {
		panic(err)
	}
	hostRating := &domain.ApartmentRating{
		ApartmentId: objID,
		//UserId:        primitive.ObjectIDFromHex(hostRatingPb.UserId),
		UserFirstName: apartmentRatingPb.UserFirstName,
		UserLastName:  apartmentRatingPb.UserLastName,
		Rating:        apartmentRatingPb.Rating,
	}

	return hostRating
}

func mapApartment(apartment *domain.Apartment) *pb.Apartment {
	apartmentPb := &pb.Apartment{
		Id:       apartment.Id.Hex(),
		Name:     apartment.Name,
		Location: apartment.Location,
		Benefits: apartment.Benefits,
>>>>>>> origin/update-proto
	}
	return apartmentPb
}

package api

import (
	"strconv"

	"github.com/NikolinaSesa/Booking/user-service/domain"
	pb "github.com/NikolinaSesa/Booking/user-service/proto"
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

func mapReservation(reservation *domain.Reservation, apartment *domain.Apartment) *pb.Reservation {
	apartmentPb := &pb.Apartment{
		Id:       apartment.Id.Hex(),
		Name:     apartment.Name,
		Location: apartment.Location,
		Benefits: apartment.Benefits,
	}
	reservationPb := &pb.Reservation{
		Id:           reservation.ID.Hex(),
		Apartment:    apartmentPb,
		StartDate:    reservation.StartDate,
		EndDate:      reservation.EndDate,
		GuestsNumber: strconv.Itoa(reservation.GuestsNumber),
	}
	return reservationPb
}

func mapFlight(flight *domain.Flight) *pb.Flight {
	flightPb := &pb.Flight{
		Id:             flight.ID.Hex(),
		Departure:      flight.Departure,
		DeparturePlace: flight.DeparturePlace,
		ArrivalPlace:   flight.ArrivalPlace,
		Price:          strconv.Itoa(flight.Price),
		NumOfFreeSeats: strconv.Itoa(flight.NumberOfFreeSeats),
	}
	return flightPb
}

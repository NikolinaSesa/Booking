package api

import (
	"context"
	"fmt"

	"github.com/NikolinaSesa/Booking/user-service/application"
	pb "github.com/NikolinaSesa/Booking/user-service/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	service *application.UserService
}

func NewUserHandler(service *application.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	fmt.Print("****************************************Tu sammm ", request.Id)
	fmt.Print("****************************************Tu sammm ", request.GetId())
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	User, err := h.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	UserPb := mapUser(User)
	response := &pb.GetResponse{
		User: UserPb,
	}

	fmt.Print("****************************************Tu sammm ", response.User.Id, response.User.FirstName)
	return response, nil
}

func (h *UserHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {

	Users, err := h.service.GetAll()
	if err != nil {
		return nil, err
	}
	var Users2 []*pb.User

	for _, user := range Users {
		UserPb := mapUser(user)
		Users2 = append(Users2, UserPb)
	}

	response := &pb.GetAllResponse{
		Users: Users2,
	}

	return response, nil
}

func (h *UserHandler) GetUserByUsernameAndPassword(ctx context.Context, loginRequest *pb.LoginRequest) (*pb.LoginResponse, error) {
	username := loginRequest.Username
	password := loginRequest.Password

	User, err := h.service.GetUserByUsernameAndPassword(username, password)
	if err != nil {
		return nil, err
	}
	UserPb := mapUser(User)
	response := &pb.LoginResponse{
		User: UserPb,
	}

	fmt.Print("****************************************Tu sammm ", response.User.FirstName, response.User.LastName)
	return response, nil
}

func (h *UserHandler) GetReservationsByGuestId(ctx context.Context, request *pb.GetReservationsByGuestIdRequest) (*pb.GetReservationsByGuestIdResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	Reservations, err := h.service.GetReservationsByGuestId(objectId)
	if err != nil {
		return nil, err
	}
	var ReseravtionsPb []*pb.Reservation

	for _, reservation := range Reservations {
		Apartment, err := h.service.GetApartment(reservation.ApartmentID)
		if err != nil {
			return nil, err
		}
		ReservationPb := mapReservation(reservation, Apartment)
		ReseravtionsPb = append(ReseravtionsPb, ReservationPb)
	}
	response := &pb.GetReservationsByGuestIdResponse{
		Reservations: ReseravtionsPb,
	}
	return response, nil
}

func (h *UserHandler) GetFlights(ctx context.Context, request *pb.GetFlightsRequest) (*pb.GetFlightsResponse, error) {
	date := request.Date

	Flights, err := h.service.GetFlightsByDeparture(date)
	if err != nil {
		return nil, err
	}
	var FlightsPb []*pb.Flight

	for _, flight := range Flights {
		FlightPb := mapFlight(flight)
		FlightsPb = append(FlightsPb, FlightPb)
	}

	response := &pb.GetFlightsResponse{
		Flights: FlightsPb,
	}
	return response, nil
}
package api

import (
	"context"
	"fmt"
	"github.com/NikolinaSesa/Booking/user-service/application"
	pb "github.com/NikolinaSesa/Booking/user-service/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
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

func (h *UserHandler) GetAllFilteredApartments(ctx context.Context, filterAllApartmentsRequest *pb.FilterAllApartmentsRequest) (*pb.FilterAllApartmentsResponse, error) {
	lowerPrice, err := strconv.Atoi(filterAllApartmentsRequest.LowerPrice)
	upperPrice, err := strconv.Atoi(filterAllApartmentsRequest.UpperPrice)
	benefit := filterAllApartmentsRequest.Benefit
	hostId, err := primitive.ObjectIDFromHex(filterAllApartmentsRequest.HostId)

	Apartments, err := h.service.GetAllFilteredApartments(lowerPrice, upperPrice, benefit, hostId)

	if err != nil {
		return nil, err
	}

	var Apartments2 []*pb.Apartment

	for _, apartment := range Apartments {
		ApartmentPb := mapApartment(apartment)
		Apartments2 = append(Apartments2, ApartmentPb)
	}

	response := &pb.FilterAllApartmentsResponse{
		Apartments: Apartments2,
	}

	return response, nil
}

func (h *UserHandler) UpdateHost(ctx context.Context, request *pb.UpdateHostRequest) (*pb.UpdateHostResponse, error) {
	hostRating := mapHostRating2(request.HostsRating)
	User, err := h.service.UpdateHost(hostRating)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateHostResponse{
		Host: mapHostRating(User),
	}, nil
}

func (h *UserHandler) UpdateApartment(ctx context.Context, request *pb.UpdateApartmentRequest) (*pb.UpdateApartmentResponse, error) {
	apartmentRating := mapApartmentRating2(request.ApartmentsRating)
	Apartment2, err := h.service.UpdateApartment(apartmentRating)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateApartmentResponse{
		Apartment: mapApartmentRating(Apartment2),
	}, nil
}

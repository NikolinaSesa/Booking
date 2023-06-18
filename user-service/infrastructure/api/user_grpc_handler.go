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

package api

import (
	"Booking/apartment-service/application"
)

type ApartmentHandler struct {
	//pb.UnimplementedUserServiceServer
	service *application.ApartmentService
}

func NewApartmentHandler(service *application.ApartmentService) *ApartmentHandler {
	return &ApartmentHandler{
		service: service,
	}
}

//
//func (h *UserHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
//	id := request.Id
//	objectId, err := primitive.ObjectIDFromHex(id)
//	if err != nil {
//		return nil, err
//	}
//	User, err := h.service.Get(objectId)
//	if err != nil {
//		return nil, err
//	}
//	UserPb := mapUser(User)
//	response := &pb.GetResponse{
//		User: UserPb,
//	}
//	return response, nil
//}

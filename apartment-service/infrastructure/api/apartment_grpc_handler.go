package api

import (
	"context"

	"Booking/apartment-service/application"
	pb "Booking/apartment-service/proto"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ApartmentHandler struct {
	pb.UnimplementedApartmentServiceServer
	service *application.ApartmentService
}

func NewApartmentHandler(service *application.ApartmentService) *ApartmentHandler {
	return &ApartmentHandler{
		service: service,
	}
}

func (h *ApartmentHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	Apartment, err := h.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	ApartmentPb := mapApartment(Apartment)
	response := &pb.GetResponse{
		Apartment: ApartmentPb,
	}
	return response, nil
}

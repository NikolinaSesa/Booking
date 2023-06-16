package api

import (
	"Booking/user-service/domain"
	pb "Booking/user-service/proto"
)

func mapUser(user *domain.User) *pb.User {
	userPb := &pb.User{
		Id:        user.Id.Hex(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role,
	}
	return userPb
}

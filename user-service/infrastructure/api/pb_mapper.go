package api

import (
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

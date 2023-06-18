package application

import (
	"github.com/NikolinaSesa/Booking/user-service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	store domain.UserStore
	//TODO
}

func NewUserService(store domain.UserStore) *UserService {
	return &UserService{
		store: store,
	}
}
func (s *UserService) Get(id primitive.ObjectID) (*domain.User, error) {
	return s.store.Get(id)
}

func (s *UserService) GetAll() ([]*domain.User, error) {
	return s.store.GetAll()
}

func (s *UserService) GetUserByUsernameAndPassword(username string, password string) (*domain.User, error) {
	return s.store.GetUserByUsernameAndPassword(username, password)
}

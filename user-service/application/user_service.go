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

func (s *UserService) UpdateHost(hostRating *domain.HostRating) (*domain.User, error) {
	return s.store.UpdateHost(hostRating)

}

func (s *UserService) GetApartment(id primitive.ObjectID) (*domain.Apartment, error) {
	return s.store.GetApartment(id)
}

func (s *UserService) GetReservation(id primitive.ObjectID) (*domain.Reservation, error) {
	return s.store.GetReservation(id)
}

func (s *UserService) GetAllApartments() ([]*domain.Apartment, error) {
	return s.store.GetAllApartments()
}

func (s *UserService) GetAllReservation() ([]*domain.Reservation, error) {
	return s.store.GetAllReservations()
}

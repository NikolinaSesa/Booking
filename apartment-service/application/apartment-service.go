package application

import (
	"github.com/NikolinaSesa/Booking/apartment-service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ApartmentService struct {
	store domain.ApartmentStore
}

func NewApartmentService(store domain.ApartmentStore) *ApartmentService {
	return &ApartmentService{
		store: store,
	}
}
func (s *ApartmentService) Get(id primitive.ObjectID) (*domain.Apartment, error) {
	return s.store.Get(id)
}

func (s *ApartmentService) GetAll() ([]*domain.Apartment, error) {
	return s.store.GetAll()
}

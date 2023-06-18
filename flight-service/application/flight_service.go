package application

import "Booking/flight-service/domain"

type FlightService struct {
	store domain.FlightStore
}

func NewFlightService(store domain.FlightStore) *FlightService {
	return &FlightService{
		store: store,
	}
}

func (s *FlightService) GetFlightsByDeparture(departure string) ([]*domain.Flight, error) {
	return s.store.GetFlightsByDeparture(departure)
}

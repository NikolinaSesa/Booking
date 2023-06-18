package domain

type FlightStore interface {
	InsertFlight(flight *Flight) error
	DeleteAllFlights()
	GetFlightsByDeparture(departure string) ([]*Flight, error)
}

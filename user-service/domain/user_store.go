package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserStore interface {
	Get(id primitive.ObjectID) (*User, error)
	GetAll() ([]*User, error)
	Insert(user *User) error
	DeleteAll()
	GetUserByUsernameAndPassword(username string, password string) (*User, error)
	UpdateHost(hostRating *HostRating) (*User, error)
	UpdateApartment(apartmentRating *ApartmentRating) (*Apartment, error)

	GetApartment(id primitive.ObjectID) (*Apartment, error)
	GetAllApartments() ([]*Apartment, error)
	InsertApartment(apartment *Apartment) error
	DeleteAllApartments()
	GetAllFilteredApartments(lowerPrice int, upperPrice int, benefit string, hostId primitive.ObjectID) ([]*Apartment, error)

	GetReservation(id primitive.ObjectID) (*Reservation, error)
	GetAllReservations() ([]*Reservation, error)
	InsertReservation(reservation *Reservation) error
	DeleteAllReservations()
}

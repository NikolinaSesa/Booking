package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserStore interface {
	Get(id primitive.ObjectID) (*User, error)
	GetAll() ([]*User, error)
	Insert(user *User) error
	DeleteAll()
	GetUserByUsernameAndPassword(username string, password string) (*User, error)
	//UpdateHostsRating(user *User) error

	GetApartment(id primitive.ObjectID) (*Apartment, error)
	GetAllApartments() ([]*Apartment, error)
	InsertApartment(apartment *Apartment) error
	DeleteAllApartments()

	GetReservation(id primitive.ObjectID) (*Reservation, error)
	GetAllReservations() ([]*Reservation, error)
	InsertReservation(reservation *Reservation) error
	DeleteAllReservations()
}

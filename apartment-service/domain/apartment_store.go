package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ApartmentStore interface {
	Get(id primitive.ObjectID) (*Apartment, error)
	GetAll() ([]*Apartment, error)
}

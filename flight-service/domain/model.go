package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Flight struct {
	ID                primitive.ObjectID `bson:"_id, omitempty"`
	Departure         string             `bson:"departure"`
	DeparturePlace    string             `bson:"departurePlace"`
	ArrivalPlace      string             `bson:"arrivalPlace"`
	Price             uint64             `bson:"price"`
	NumberOfFreeSeats uint64             `bson:"numberOfFreeSeats"`
}

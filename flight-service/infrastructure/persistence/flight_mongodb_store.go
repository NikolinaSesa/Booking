package persistence

import (
	"Booking/flight-service/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "booking"
	COLLECTION = "flights"
)

type FlightMongoDBStore struct {
	flights *mongo.Collection
}

func NewFlightMongoDBStore(client *mongo.Client) domain.FlightStore {
	flights := client.Database(DATABASE).Collection(COLLECTION)
	return &FlightMongoDBStore{
		flights: flights,
	}
}

func (s *FlightMongoDBStore) GetFlightsByDeparture(departure string) ([]*domain.Flight, error) {
	filter := bson.M{"departure": departure}
	return s.filterFlights(filter)
}

func (s *FlightMongoDBStore) InsertFlight(flight *domain.Flight) error {
	result, err := s.flights.InsertOne(context.TODO(), flight)
	if err != nil {
		return err
	}
	flight.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (s *FlightMongoDBStore) DeleteAllFlights() {
	s.flights.DeleteMany(context.TODO(), bson.D{{}})
}

func (s *FlightMongoDBStore) filterFlights(filter interface{}) ([]*domain.Flight, error) {
	cursor, err := s.flights.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decodeFlights(cursor)
}

func decodeFlights(cursor *mongo.Cursor) (flights []*domain.Flight, err error) {
	for cursor.Next(context.TODO()) {
		var Flight domain.Flight
		err = cursor.Decode(&Flight)
		if err != nil {
			return
		}
		flights = append(flights, &Flight)
	}
	err = cursor.Err()
	return
}

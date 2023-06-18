package persistence

import (
	"context"

	"github.com/NikolinaSesa/Booking/apartment-service/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "booking"
	COLLECTION = "apartments"
)

type ApartmentMongoDBStore struct {
	apartments *mongo.Collection
}

func NewApartmentMongoDBStore(client *mongo.Client) domain.ApartmentStore {
	apartments := client.Database(DATABASE).Collection(COLLECTION)
	return &ApartmentMongoDBStore{
		apartments: apartments,
	}
}

func (s *ApartmentMongoDBStore) Get(id primitive.ObjectID) (*domain.Apartment, error) {
	filter := bson.M{"_id": id}
	return s.filterOne(filter)
}

func (s *ApartmentMongoDBStore) GetAll() ([]*domain.Apartment, error) {
	filter := bson.D{{}}
	return s.filter(filter)
}

func (s *ApartmentMongoDBStore) Insert(apartment *domain.Apartment) error {
	result, err := s.apartments.InsertOne(context.TODO(), apartment)
	if err != nil {
		return err
	}
	apartment.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (s *ApartmentMongoDBStore) DeleteAll() {
	s.apartments.DeleteMany(context.TODO(), bson.D{{}})
}

func (s *ApartmentMongoDBStore) filter(filter interface{}) ([]*domain.Apartment, error) {
	cursor, err := s.apartments.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (s *ApartmentMongoDBStore) filterOne(filter interface{}) (Apartment *domain.Apartment, err error) {
	result := s.apartments.FindOne(context.TODO(), filter)
	err = result.Decode(&Apartment)
	return
}

func decode(cursor *mongo.Cursor) (apartments []*domain.Apartment, err error) {
	for cursor.Next(context.TODO()) {
		var Apartment domain.Apartment
		err = cursor.Decode(&Apartment)
		if err != nil {
			return
		}
		apartments = append(apartments, &Apartment)
	}
	err = cursor.Err()
	return
}

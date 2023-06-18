package persistence

import (
	"context"

	"github.com/NikolinaSesa/Booking/user-service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE    = "booking"
	COLLECTION1 = "users"
	COLLECTION2 = "apartments"
	COLLECTION3 = "reservations"
	COLLECTION4 = "flights"
)

type UserMongoDBStore struct {
	users        *mongo.Collection
	apartments   *mongo.Collection
	reservations *mongo.Collection
	flights      *mongo.Collection
}

func NewUserMongoDBStore(client *mongo.Client) domain.UserStore {
	users := client.Database(DATABASE).Collection(COLLECTION1)
	apartments := client.Database(DATABASE).Collection(COLLECTION2)
	reservations := client.Database(DATABASE).Collection(COLLECTION3)
	flights := client.Database(DATABASE).Collection(COLLECTION4)
	return &UserMongoDBStore{
		users:        users,
		apartments:   apartments,
		reservations: reservations,
		flights:      flights,
	}
}

func (s *UserMongoDBStore) Get(id primitive.ObjectID) (*domain.User, error) {
	filter := bson.M{"_id": id}
	return s.filterOne(filter)
}

func (s *UserMongoDBStore) GetApartment(id primitive.ObjectID) (*domain.Apartment, error) {
	filter := bson.M{"_id": id}
	return s.filterOneApartment(filter)
}

func (s *UserMongoDBStore) GetReservation(id primitive.ObjectID) (*domain.Reservation, error) {
	filter := bson.M{"_id": id}
	return s.filterOneReservation(filter)
}

func (s *UserMongoDBStore) GetUserByUsernameAndPassword(username string, password string) (*domain.User, error) {
	filter := bson.M{"username": username, "password": password}
	return s.filterOne(filter)
}

func (s *UserMongoDBStore) GetAll() ([]*domain.User, error) {
	filter := bson.D{{}}
	return s.filter(filter)
}

func (s *UserMongoDBStore) GetAllApartments() ([]*domain.Apartment, error) {
	filter := bson.D{{}}
	return s.filterApartments(filter)
}

func (s *UserMongoDBStore) GetAllReservations() ([]*domain.Reservation, error) {
	filter := bson.D{{}}
	return s.filterReservations(filter)
}

func (s *UserMongoDBStore) Insert(user *domain.User) error {
	result, err := s.users.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	user.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (s *UserMongoDBStore) InsertApartment(apartment *domain.Apartment) error {
	result, err := s.apartments.InsertOne(context.TODO(), apartment)
	if err != nil {
		return err
	}
	apartment.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}
func (s *UserMongoDBStore) InsertReservation(reservation *domain.Reservation) error {
	result, err := s.reservations.InsertOne(context.TODO(), reservation)
	if err != nil {
		return err
	}
	reservation.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (s *UserMongoDBStore) DeleteAll() {
	s.users.DeleteMany(context.TODO(), bson.D{{}})
}

func (s *UserMongoDBStore) DeleteAllApartments() {
	s.apartments.DeleteMany(context.TODO(), bson.D{{}})
}

func (s *UserMongoDBStore) DeleteAllReservations() {
	s.reservations.DeleteMany(context.TODO(), bson.D{{}})
}

func (s *UserMongoDBStore) GetReservationsByGuestId(id primitive.ObjectID) ([]*domain.Reservation, error) {
	filter := bson.M{"guestId": id}
	return s.filterReservations(filter)
}

func (s *UserMongoDBStore) GetFlightsByDeparture(departure string) ([]*domain.Flight, error) {
	filter := bson.M{"departure": departure}
	return s.filterFlights(filter)
}
func (s *UserMongoDBStore) InsertFlight(flight *domain.Flight) error {
	result, err := s.flights.InsertOne(context.TODO(), flight)
	if err != nil {
		return err
	}
	flight.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (s *UserMongoDBStore) DeleteAllFlights() {
	s.flights.DeleteMany(context.TODO(), bson.D{{}})
}

func (s *UserMongoDBStore) filter(filter interface{}) ([]*domain.User, error) {
	cursor, err := s.users.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (s *UserMongoDBStore) filterApartments(filter interface{}) ([]*domain.Apartment, error) {
	cursor, err := s.apartments.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decodeApartment(cursor)
}

func (s *UserMongoDBStore) filterReservations(filter interface{}) ([]*domain.Reservation, error) {
	cursor, err := s.reservations.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decodeReservation(cursor)
}

func (s *UserMongoDBStore) filterFlights(filter interface{}) ([]*domain.Flight, error) {
	cursor, err := s.flights.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decodeFlights(cursor)
}

func (s *UserMongoDBStore) filterOne(filter interface{}) (User *domain.User, err error) {
	result := s.users.FindOne(context.TODO(), filter)
	err = result.Decode(&User)
	return
}

func (s *UserMongoDBStore) filterOneApartment(filter interface{}) (Apartment *domain.Apartment, err error) {
	result := s.apartments.FindOne(context.TODO(), filter)
	err = result.Decode(&Apartment)
	return
}

func (s *UserMongoDBStore) filterOneReservation(filter interface{}) (Reservation *domain.Reservation, err error) {
	result := s.reservations.FindOne(context.TODO(), filter)
	err = result.Decode(&Reservation)
	return
}

func decode(cursor *mongo.Cursor) (users []*domain.User, err error) {
	for cursor.Next(context.TODO()) {
		var User domain.User
		err = cursor.Decode(&User)
		if err != nil {
			return
		}
		users = append(users, &User)
	}
	err = cursor.Err()
	return
}

func decodeApartment(cursor *mongo.Cursor) (apartments []*domain.Apartment, err error) {
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

func decodeReservation(cursor *mongo.Cursor) (reservations []*domain.Reservation, err error) {
	for cursor.Next(context.TODO()) {
		var Reservation domain.Reservation
		err = cursor.Decode(&Reservation)
		if err != nil {
			return
		}
		reservations = append(reservations, &Reservation)
	}
	err = cursor.Err()
	return
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

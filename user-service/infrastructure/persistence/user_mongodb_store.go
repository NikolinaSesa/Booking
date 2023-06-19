package persistence

import (
	"context"
	"errors"
	"fmt"
	"github.com/NikolinaSesa/Booking/user-service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
	"time"
)

const (
	DATABASE    = "booking"
	COLLECTION1 = "users"
	COLLECTION2 = "apartments"
	COLLECTION3 = "reservations"
)

type UserMongoDBStore struct {
	users        *mongo.Collection
	apartments   *mongo.Collection
	reservations *mongo.Collection
}

func NewUserMongoDBStore(client *mongo.Client) domain.UserStore {
	users := client.Database(DATABASE).Collection(COLLECTION1)
	apartments := client.Database(DATABASE).Collection(COLLECTION2)
	reservations := client.Database(DATABASE).Collection(COLLECTION3)
	return &UserMongoDBStore{
		users:        users,
		apartments:   apartments,
		reservations: reservations,
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

func (s *UserMongoDBStore) UpdateHost(hostRating *domain.HostRating) (*domain.User, error) {
	fmt.Println("################################ ", hostRating.UserFirstName, " #######################")

	filter := bson.M{"_id": hostRating.HostId}

	host, _ := s.filterOne(filter)

	var structa domain.Rating
	structa.GuestFirstName = hostRating.UserFirstName
	structa.GuestLastName = hostRating.UserLastName
	structa.RatedAt = time.Now()

	number, _ := strconv.ParseInt(hostRating.Rating, 10, 64)

	structa.Rating = int(number)

	host.Ratings = append(host.Ratings, structa)

	var avg = 0.0

	for i := 0; i < len(host.Ratings); i++ {
		var rating float64 = float64(host.Ratings[i].Rating)
		avg += rating
	}

	avg = avg / (float64(len(host.Ratings)))

	result, err := s.users.UpdateOne(
		context.TODO(),
		bson.M{"_id": hostRating.HostId},
		bson.D{
			{"$set", bson.D{{"firstName", hostRating.UserFirstName}}},
			{"$set", bson.D{{"ratings", host.Ratings}}},
			{"$set", bson.D{{"avgRating", avg}}},
		},
	)
	if err != nil {
		return nil, err
	}

	if result.MatchedCount != 1 {
		return nil, errors.New("one document should've been updated")
	}

	filter2 := bson.M{"_id": hostRating.HostId}
	return s.filterOne(filter2)
}

func (s *UserMongoDBStore) UpdateApartment(apartmentRating *domain.ApartmentRating) (*domain.Apartment, error) {
	fmt.Println("################################ ", apartmentRating.UserFirstName, " #######################")

	filter := bson.M{"_id": apartmentRating.ApartmentId}

	apartment, _ := s.filterOneApartment(filter)

	var structa domain.RatingApartment
	structa.GuestFirstName = apartmentRating.UserFirstName
	structa.GuestLastName = apartmentRating.UserLastName
	structa.RatedAt = time.Now()

	number, _ := strconv.ParseInt(apartmentRating.Rating, 10, 64)

	structa.Rating = int(number)

	apartment.Ratings = append(apartment.Ratings, structa)

	var avg = 0.0

	for i := 0; i < len(apartment.Ratings); i++ {
		var rating float64 = float64(apartment.Ratings[i].Rating)
		avg += rating
	}

	avg = avg / (float64(len(apartment.Ratings)))

	result, err := s.apartments.UpdateOne(
		context.TODO(),
		bson.M{"_id": apartmentRating.ApartmentId},
		bson.D{
			//{"$set", bson.D{{"name", apartmentRating.UserFirstName}}},
			{"$set", bson.D{{"ratings", apartment.Ratings}}},
			{"$set", bson.D{{"avgRating", avg}}},
		},
	)
	if err != nil {
		return nil, err
	}

	if result.MatchedCount != 1 {
		return nil, errors.New("one document should've been updated")
	}

	filter2 := bson.M{"_id": apartmentRating.ApartmentId}
	return s.filterOneApartment(filter2)
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

func (s *UserMongoDBStore) GetAllFilteredApartments(lowerPrice int, upperPrice int, benefit string, hostId primitive.ObjectID) ([]*domain.Apartment, error) {
	fmt.Println("****************************************EEEE ", lowerPrice, upperPrice, benefit, hostId)
	filter := bson.M{"hostId": hostId, "benefits": benefit}
	var apartments []*domain.Apartment
	apartments, _ = s.filterApartments(filter)
	if apartments != nil {

	}
	var apartmentsToSend []*domain.Apartment
	for i := 0; i < len(apartments); i++ {
		number, _ := strconv.ParseInt(apartments[i].GeneralPrice, 10, 64)
		fmt.Println("################################### ", number, " ###################################")
		fmt.Println("################################### nepotrebna provera:  ", apartments[i].GeneralPrice, " ###################################")
		if int(number) >= lowerPrice && int(number) <= upperPrice {
			apartmentsToSend = append(apartmentsToSend, apartments[i])
		}
	}
	//filter := bson.M{"hostId": hostId, "benefits": benefit, "generalPrice": bson.M{"$gte": lowerPrice}}
	//filter := bson.M{"hostId": hostId, "benefits": benefit, "generalPrice": lowerPrice}
	return apartmentsToSend, nil
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

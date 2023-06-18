package persistence

import (
	"context"

	"github.com/NikolinaSesa/Booking/user-service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "booking"
	COLLECTION = "users"
)

type UserMongoDBStore struct {
	users *mongo.Collection
}

func NewUserMongoDBStore(client *mongo.Client) domain.UserStore {
	users := client.Database(DATABASE).Collection(COLLECTION)
	return &UserMongoDBStore{
		users: users,
	}
}

func (s *UserMongoDBStore) Get(id primitive.ObjectID) (*domain.User, error) {
	filter := bson.M{"_id": id}
	return s.filterOne(filter)
}

func (s *UserMongoDBStore) GetUserByUsernameAndPassword(username string, password string) (*domain.User, error) {
	filter := bson.M{"username": username, "password": password}
	return s.filterOne(filter)
}

func (s *UserMongoDBStore) GetAll() ([]*domain.User, error) {
	filter := bson.D{{}}
	return s.filter(filter)
}

func (s *UserMongoDBStore) Insert(user *domain.User) error {
	result, err := s.users.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	user.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (s *UserMongoDBStore) DeleteAll() {
	s.users.DeleteMany(context.TODO(), bson.D{{}})
}

func (s *UserMongoDBStore) filter(filter interface{}) ([]*domain.User, error) {
	cursor, err := s.users.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (s *UserMongoDBStore) filterOne(filter interface{}) (User *domain.User, err error) {
	result := s.users.FindOne(context.TODO(), filter)
	err = result.Decode(&User)
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

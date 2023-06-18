package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Guest struct {
	Id        primitive.ObjectID `bson:"id"`
	FirstName string             `bson:"firstName"`
	LastName  string             `bson:"lastName"`
}

type Rating struct {
	Guest   Guest     `bson:"guest"`
	RatedAt time.Time `bson:"ratedAt"`
	Rating  int       `bson:"rating"`
}

type User struct {
	Id        primitive.ObjectID `bson:"_id"`
	FirstName string             `bson:"firstName" json:"firstName"`
	LastName  string             `bson:"lastName" json:"lastName"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
	Username  string             `bson:"username" json:"username"`
	Address   string             `bson:"address" json:"address"`
	Role      string             `bson:"role" json:"role"`
	Ratings   []Rating           `bson:"ratings"`
	AvgRating float64            `bson:"avgRating"`
	Mark      bool               `bson:"mark"`
}

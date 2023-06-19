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
	Guest          Guest     `bson:"guest"`
	RatedAt        time.Time `bson:"ratedAt"`
	Rating         int       `bson:"rating"`
	GuestFirstName string    `bson:"guestFirstName"`
	GuestLastName  string    `bson:"guestLastName"`
}

type RatingApartment struct {
	RatedAt        time.Time `bson:"ratedAt"`
	Rating         int       `bson:"rating"`
	GuestFirstName string    `bson:"guestFirstName"`
	GuestLastName  string    `bson:"guestLastName"`
}

type NotificationType byte

const (
	NewRating NotificationType = iota
	NewApartmentRating
)

type Notification struct {
	Message string           `bson:"message"`
	SendAt  time.Time        `bson:"sendAt"`
	On      bool             `bson:"on"`
	Type    NotificationType `bson:"type"`
}

type User struct {
	Id            primitive.ObjectID `bson:"_id"`
	FirstName     string             `bson:"firstName" json:"firstName"`
	LastName      string             `bson:"lastName" json:"lastName"`
	Email         string             `bson:"email" json:"email"`
	Password      string             `bson:"password" json:"password"`
	Username      string             `bson:"username" json:"username"`
	Address       string             `bson:"address" json:"address"`
	Role          string             `bson:"role" json:"role"`
	Ratings       []Rating           `bson:"ratings"`
	AvgRating     float64            `bson:"avgRating"`
	Mark          string             `bson:"mark"`
	Notifications []Notification     `bson:"notifications"`
}

type Apartment struct {
	Id                   primitive.ObjectID `bson:"_id" json:"id"`
	HostId               primitive.ObjectID `bson:"hostId" json:"hostId"`
	Name                 string             `bson:"name" json:"name"`
	Location             string             `bson:"location" json:"location"`
	Benefits             string             `bson:"benefits" json:"benefits"`
	MinGuestsNumber      int                `bson:"minGuestsNumber" json:"minGuestsNumber"`
	MaxGuestsNumber      int                `bson:"maxGuestsNumber" json:"maxGuestsNumber"`
	AutomaticReservation bool               `bson:"automaticReservation" json:"automaticReservation"`
	PriceList            []*PriceListItem   `bson:"pricelist" json:"pricelist"`
	GeneralPrice         string             `bson:"generalPrice" json:"generalPrice"`
	Ratings              []RatingApartment  `bson:"ratings"`
}

type AvailableApartment struct {
	Apartment  *Apartment
	TotalPrice int
	UnitPrice  int
}

type PriceListItem struct {
	AvailabilityStartDate string `bson:"availabilityStartDate" json:"availabilityStartDate"`
	AvailabilityEndDate   string `bson:"availabilityEndDate" json:"availabilityEndDate"`
	Price                 int    `bson:"price" json:"price"`
	UnitPrice             int    `bson:"unitPrice" json:"unitPrice"`
}

type Reservation struct {
	ID           primitive.ObjectID `bson:"_id"`
	GuestID      primitive.ObjectID `bson:"guestId"`
	ApartmentID  primitive.ObjectID `bson:"apartmentId"`
	StartDate    string             `bson:"startDate"`
	EndDate      string             `bson:"endDate"`
	GuestsNumber int                `bson:"guestsNumber"`
}

type HostRating struct {
	HostId        primitive.ObjectID `bson:"_id"`
	UserId        primitive.ObjectID `bson:"userId"`
	UserFirstName string             `bson:"userFirstName"`
	UserLastName  string             `bson:"userLastName"`
	Rating        string             `bson:"rating"`
}

type ApartmentRating struct {
	ApartmentId   primitive.ObjectID `bson:"_id"`
	UserId        primitive.ObjectID `bson:"userId"`
	UserFirstName string             `bson:"userFirstName"`
	UserLastName  string             `bson:"userLastName"`
	Rating        string             `bson:"rating"`
}

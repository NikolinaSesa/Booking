package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Rating struct {
	RatedAt        time.Time `bson:"ratedAt"`
	Rating         int       `bson:"rating"`
	GuestFirstName string    `bson:"guestFirstName"`
	GuestLastName  string    `bson:"guestLastName"`
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
	Ratings              []Rating           `bson:"ratings"`
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

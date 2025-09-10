package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type TripModel struct {
	ID primitive.ObjectID `json:"id"`
}

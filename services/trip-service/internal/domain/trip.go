package domain

import (
	"context"
	"ride-sharing/shared/types"

	tripType "ride-sharing/services/trip-service/pkg/types"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TripModel struct {
	ID       primitive.ObjectID `json:"id"`
	UserID   string             `json:"user_id"`
	Status   string             `json:"status"`
	RideFare *RideFareModel     `json:"ride_fare"`
}

type TripRepository interface {
	CreateTrip(ctx context.Context, trip *TripModel) (*TripModel, error)
}

type TripService interface {
	CreateTrip(ctx context.Context, fare *RideFareModel) (*TripModel, error)
	GetRoute(ctx context.Context, pickup *types.Coordinate, destination *types.Coordinate) (*tripType.OsrmApiResponse, error)
}

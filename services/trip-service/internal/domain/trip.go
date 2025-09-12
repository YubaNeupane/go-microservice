package domain

import (
	"context"
	"ride-sharing/shared/types"

	tripType "ride-sharing/services/trip-service/pkg/types"

	pb "ride-sharing/shared/proto/trip"

	tripTypes "ride-sharing/services/trip-service/pkg/types"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TripModel struct {
	ID       primitive.ObjectID `json:"id"`
	UserID   string             `json:"user_id"`
	Status   string             `json:"status"`
	RideFare *RideFareModel     `json:"ride_fare"`
	Driver   *pb.TripDriver     `json:"driver"`
}

type TripRepository interface {
	CreateTrip(ctx context.Context, trip *TripModel) (*TripModel, error)
	SaveRideFare(ctx context.Context, fare *RideFareModel) error

	GetRideFareByID(ctx context.Context, id string) (*RideFareModel, error)
}

type TripService interface {
	CreateTrip(ctx context.Context, fare *RideFareModel) (*TripModel, error)
	GetRoute(ctx context.Context, pickup *types.Coordinate, destination *types.Coordinate) (*tripType.OsrmApiResponse, error)
	EstimatePackagesPriceWithRoute(route *tripType.OsrmApiResponse) []*RideFareModel
	GenerateTripFares(ctx context.Context, fares []*RideFareModel, userID string, route *tripTypes.OsrmApiResponse) ([]*RideFareModel, error)

	GetAndValidateFare(ctx context.Context, fareId string, userID string) (*RideFareModel, error)
}

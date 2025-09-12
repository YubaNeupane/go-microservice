package domain

import (
	pb "ride-sharing/shared/proto/trip"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RideFareModel struct {
	ID                primitive.ObjectID `json:"id"`
	UserID            string             `json:"user_id"`
	PackageSlug       string             `json:"package_slug"` //e.g van, luxury, sedan
	TotalPriceInCents float64            `json:"total_price_in_cents"`
}

func (r *RideFareModel) ToProto() *pb.RideFare {
	return &pb.RideFare{
		Id:                r.ID.Hex(),
		UserID:            r.UserID,
		PackageSlug:       r.PackageSlug,
		TotalPriceInCents: r.TotalPriceInCents,
	}

}

func ToRideFareProto(fares []*RideFareModel) []*pb.RideFare {

	var protoFares []*pb.RideFare
	for _, q := range fares {
		protoFares = append(protoFares, q.ToProto())
	}

	return protoFares
}

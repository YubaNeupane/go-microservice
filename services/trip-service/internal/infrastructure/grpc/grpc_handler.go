package grpc

import (
	"context"
	"log"
	"ride-sharing/services/trip-service/internal/domain"
	pb "ride-sharing/shared/proto/trip"
	"ride-sharing/shared/types"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type gRPCHandler struct {
	pb.UnimplementedTripServiceServer
	service domain.TripService
}

func NewGRPCHandler(server *grpc.Server, service domain.TripService) *gRPCHandler {
	handler := &gRPCHandler{
		service: service,
	}

	pb.RegisterTripServiceServer(server, handler)
	return handler
}

func (h *gRPCHandler) PreviewTrip(ctx context.Context, previewTripRequest *pb.PreviewTripRequest) (*pb.PreviewTripResponse, error) {

	pickup := previewTripRequest.GetStartLocation()
	destination := previewTripRequest.GetEndLocation()

	pickUpCoord := &types.Coordinate{
		Latitude:  pickup.Latitude,
		Longitude: pickup.Longitude,
	}

	destinationCoord := &types.Coordinate{
		Latitude:  destination.Latitude,
		Longitude: destination.Longitude,
	}

	route, err := h.service.GetRoute(ctx, pickUpCoord, destinationCoord)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, "Failed to get route: %v", err)
	}

	estimatedFares := h.service.EstimatePackagesPriceWithRoute(route)
	fares, err := h.service.GenerateTripFares(ctx, estimatedFares, previewTripRequest.UserID, route)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, "Failed to generate ride fares: %v", err)
	}

	return &pb.PreviewTripResponse{
		Route:     route.ToProto(),
		RideFares: domain.ToRideFareProto(fares),
	}, nil

}

func (h *gRPCHandler) CreateTrip(ctx context.Context, createTripRequest *pb.CreateTripRequest) (*pb.CreateTripResponse, error) {

	fare, err := h.service.GetAndValidateFare(ctx, createTripRequest.RideFareID, createTripRequest.UserID)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, "Failed to GetAndValidateFare: %v", err)
	}

	trip, err := h.service.CreateTrip(ctx, fare)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, "Failed to CreateTrip: %v", err)
	}

	//Publish to event on async common module

	return &pb.CreateTripResponse{
		TripID: trip.ID.Hex(),
	}, nil

}

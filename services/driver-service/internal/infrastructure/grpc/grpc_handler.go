package grpc

import (
	"context"
	"ride-sharing/services/driver-service/internal/service"

	pb "ride-sharing/shared/proto/driver"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcHandler struct {
	pb.UnimplementedDriverServiceServer
	*service.Service
}

func NewGrpcHandler(s *grpc.Server, service *service.Service) {
	handler := &grpcHandler{}
	pb.RegisterDriverServiceServer(s, handler)
}

func (h *grpcHandler) RegisterDriver(ctx context.Context, registerDriverRequest *pb.RegisterDriverRequest) (*pb.RegisterDriverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterDriver not implemented")
}
func (h *grpcHandler) UnregisterDriver(ctx context.Context, registerDriverRequest *pb.RegisterDriverRequest) (*pb.RegisterDriverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnregisterDriver not implemented")
}

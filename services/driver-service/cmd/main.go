package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"ride-sharing/services/driver-service/internal/infrastructure/grpc"
	"ride-sharing/services/driver-service/internal/service"
	"ride-sharing/shared/env"
	"syscall"

	grpcserver "google.golang.org/grpc"
)

var (
	GrpcAddr = env.GetString("GRPC_ADDR", ":9092")
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
		<-sigCh
		cancel()
	}()

	lis, err := net.Listen("tcp", GrpcAddr)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	service := service.NewService()

	//STARTING the grpc Server
	grpcServer := grpcserver.NewServer()
	grpc.NewGrpcHandler(grpcServer, service)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Printf("failed to serve %v", err)
			cancel()
		}
	}()

	log.Printf("Starting gRPC server Driver service on port %s", lis.Addr().String())
	// wait for the shutdown signal
	<-ctx.Done()
	log.Printf("Shutting down the server %v", err)
	grpcServer.GracefulStop()

}

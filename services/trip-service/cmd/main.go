package main

import (
	"context"
	"fmt"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"time"
)

func main() {
	ctx := context.Background()
	inmemRepo := repository.NewInmemRepository()
	svc := service.NewService(inmemRepo)

	fmt.Println("Trip Service is running...")
	fare := &domain.RideFareModel{
		UserID:            "user123",
		TotalPriceInCents: 25.50,
	}
	svc.CreateTrip(ctx, fare)

	// Keep the service running for now
	for {
		time.Sleep(1 * time.Second)
	}

}

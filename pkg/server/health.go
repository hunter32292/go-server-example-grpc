package server

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type HealthChecker struct{}

func NewHealthChecker() *HealthChecker {
	return &HealthChecker{}
}

func (s *HealthChecker) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {

	log.Println("Serving the Check request for health check")

	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

func (s *HealthChecker) Watch(req *grpc_health_v1.HealthCheckRequest, server grpc_health_v1.Health_WatchServer) error {
	log.Println("Serving the Watch request for health check")
	return server.Send(&grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	})
}

func RegisterHC(s *grpc.Server) {
	healthService := NewHealthChecker()
	grpc_health_v1.RegisterHealthServer(s, healthService)
}

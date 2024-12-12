package env

import (
	"github.com/sergei-galichev/project-console-chat/auth/internal/config"
	"log/slog"
	"net"
	"os"
)

const (
	grpcHostEnvName = "GRPC_HOST"
	grpcPortEnvName = "GRPC_PORT"

	grpcHostDefault = "localhost"
	grpcPortDefault = "50051"
)

type grpcConfig struct {
	host string
	port string
}

// NewGRPCConfig returns gRPC config
func NewGRPCConfig() config.GRPCConfig {
	host := os.Getenv(grpcHostEnvName)
	if len(host) == 0 {
		slog.Info("env 'GRPC_HOST' not found. used default")

		host = grpcHostDefault
	}

	port := os.Getenv(grpcPortEnvName)
	if len(port) == 0 {
		slog.Info("env 'GRPC_PORT' not found. used default")

		port = grpcPortDefault
	}

	return &grpcConfig{
		host: host,
		port: port,
	}
}

func (cfg *grpcConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}

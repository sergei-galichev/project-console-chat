package main

import (
	"context"
	"log"
	"net"

	"github.com/sergei-galichev/project-console-chat/auth/internal/config"
	"github.com/sergei-galichev/project-console-chat/auth/internal/config/env"

	"github.com/brianvoe/gofakeit"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/sergei-galichev/project-console-chat/auth/pkg/user_v1"
)

func init() {
	if err := config.LoadConfig("config/local.env"); err != nil {
		panic("Error loading .env file")
	}
}

type server struct {
	pb.UnimplementedUserV1Server
}

func (s *server) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	_, _ = ctx, req
	return &pb.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	var role pb.Role

	id := req.GetId()

	if id == 1 {
		role = pb.Role_ADMIN
	} else {
		role = pb.Role_USER
	}

	return &pb.GetResponse{
		Id:        id,
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Role:      role,
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}, nil
}

func (s *server) Update(ctx context.Context, req *pb.UpdateRequest) (*empty.Empty, error) {
	_, _ = ctx, req
	return &empty.Empty{}, nil
}

func (s *server) Delete(ctx context.Context, req *pb.DeleteRequest) (*empty.Empty, error) {
	_, _ = ctx, req
	return &empty.Empty{}, nil
}

func main() {
	grpcCfg := env.NewGRPCConfig()

	lis, err := net.Listen("tcp", grpcCfg.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	reflection.Register(s)

	pb.RegisterUserV1Server(s, &server{})

	log.Println("starting auth server on address: ", grpcCfg.Address())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

package main

import (
	"auth/internal/config"
	"auth/internal/config/env"
	user_pb "auth/pkg/user_v1"
	"context"

	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func init() {
	if err := config.LoadConfig(".env"); err != nil {
		panic("Error loading .env file")
	}
}

type server struct {
	user_pb.UnimplementedUserV1Server
}

func (s *server) Create(ctx context.Context, req *user_pb.CreateRequest) (*user_pb.CreateResponse, error) {
	_, _ = ctx, req
	return &user_pb.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) Get(ctx context.Context, req *user_pb.GetRequest) (*user_pb.GetResponse, error) {
	_ = ctx
	id := req.GetId()

	return &user_pb.GetResponse{
		User: &user_pb.User{
			Id: id,
			Info: &user_pb.UserInfo{
				Name:      gofakeit.Name(),
				Email:     gofakeit.Email(),
				Role:      user_pb.Role_USER,
				CreatedAt: timestamppb.New(gofakeit.Date()),
				UpdatedAt: timestamppb.New(gofakeit.Date()),
			},
		},
	}, nil
}

func (s *server) Update(ctx context.Context, req *user_pb.UpdateRequest) (*empty.Empty, error) {
	_, _ = ctx, req
	return &empty.Empty{}, nil
}

func (s *server) Delete(ctx context.Context, req *user_pb.DeleteRequest) (*empty.Empty, error) {
	_, _ = ctx, req
	return &empty.Empty{}, nil
}

func main() {
	grpcCfg, err := env.NewGRPCConfig()
	if err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", grpcCfg.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	user_pb.RegisterUserV1Server(s, &server{})

	log.Println("starting auth server on address: ", grpcCfg.Address())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

package main

import (
	"context"
	"log"
	"net"

	"github.com/sergei-galichev/project-console-chat/chat-server/internal/config"
	"github.com/sergei-galichev/project-console-chat/chat-server/internal/config/env"

	"github.com/brianvoe/gofakeit"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/sergei-galichev/project-console-chat/chat-server/pkg/chat_v1"
)

func init() {
	if err := config.LoadConfig("config/local.env"); err != nil {
		panic("Error loading .env file")
	}
}

type server struct {
	pb.UnimplementedChatV1Server
}

func (s *server) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	_, _ = ctx, req
	return &pb.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) Delete(ctx context.Context, req *pb.DeleteRequest) (*empty.Empty, error) {
	_, _ = ctx, req
	return &empty.Empty{}, nil
}

func (s *server) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*empty.Empty, error) {
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

	pb.RegisterChatV1Server(s, &server{})

	log.Println("starting chat server on address: ", grpcCfg.Address())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

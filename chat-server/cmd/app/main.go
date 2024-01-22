package main

import (
	"chat-server/internal/config"
	"chat-server/internal/config/env"
	chat_pb "chat-server/pkg/chat_v1"
	"context"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func init() {
	if err := config.LoadConfig(".env"); err != nil {
		panic("Error loading .env file")
	}
}

type server struct {
	chat_pb.UnimplementedChatV1Server
}

func (s *server) Create(ctx context.Context, req *chat_pb.CreateRequest) (*chat_pb.CreateResponse, error) {
	_, _ = ctx, req
	return &chat_pb.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) Delete(ctx context.Context, req *chat_pb.DeleteRequest) (*empty.Empty, error) {
	_, _ = ctx, req
	return &empty.Empty{}, nil
}

func (s *server) SendMessage(ctx context.Context, req *chat_pb.SendMessageRequest) (*empty.Empty, error) {
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
	chat_pb.RegisterChatV1Server(s, &server{})

	log.Println("starting chat server on address: ", grpcCfg.Address())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

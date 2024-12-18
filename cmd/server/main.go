package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"file-storage/pkg/fileservicepb" 
	"file-storage/internal/delivery"
	"file-storage/internal/repository"
	"file-storage/internal/usecase"
)

func main() {
	repo := repository.NewFileRepo("./storage")
	service := usecase.NewFileService(repo)
	handler := delivery.NewGRPCHandler(service)

	grpcServer := grpc.NewServer()
	fileservicepb.RegisterFileServiceServer(grpcServer, handler)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("Server is running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}


package main

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	"google.golang.org/grpc/reflection"

	"multiplayer-modes-service/internal/handlers"
	pb "multiplayer-modes-service/internal/models"
)

func main() {
	// Load environment Variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT must be set")
	}

	// Initialize gRPC server (placeholder for now)
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}


	grpcServer := grpc.NewServer()

	// Register Services
	server := handlers.NewMultiplayerServiceServer()
	pb.RegisterMultiplayerServiceServer(grpcServer, server)

	// Enable reflection
  reflection.Register(grpcServer)

	log.Printf("Server listening on port %s", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}
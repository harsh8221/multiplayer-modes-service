package handlers 

import (
	"context"

	pb "multiplayer-modes-service/internal/models"
)

type MultiplayerServiceServer struct {
	pb.UnimplementedMultiplayerServiceServer
}

func (s *MultiplayerServiceServer) GetPopularModes(ctx context.Context, req *pb.PopularModesRequest) (*pb.PopularModesResponse, error) {
	// TODO: Implement logic to get popular modes

	return &pb.PopularModesResponse{}, nil
}

func (s *MultiplayerServiceServer) ReportModePlaying(ctx context.Context, req *pb.ModePlayingRequest) (*pb.ModePlayingResponse, error) {
	// TODO: Implement logic to report mode playing
	return &pb.ModePlayingResponse{}, nil
}
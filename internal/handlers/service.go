package handlers 

import (
	"context"

	pb "multiplayer-modes-service/internal/models"
	"multiplayer-modes-service/internal/business"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MultiplayerServiceServer struct {
	pb.UnimplementedMultiplayerServiceServer
	logic *business.BusinessLogic
}

func NewMultiplayerServiceServer() *MultiplayerServiceServer {
	return &MultiplayerServiceServer{
		logic: business.NewBusinessLogic(),
	}
}

func (s *MultiplayerServiceServer) GetPopularModes(ctx context.Context, req *pb.PopularModesRequest) (*pb.PopularModesResponse, error) {
	
	areaCode := req.AreaCode

	modes, err := s.logic.GetPopularModes(ctx, areaCode)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get popular modes: %v", err)
	}

	return &pb.PopularModesResponse{
		Modes: modes,
	}, nil
}

func (s *MultiplayerServiceServer) ReportModePlaying(ctx context.Context, req *pb.ModePlayingRequest) (*pb.ModePlayingResponse, error) {

	areaCode := req.AreaCode
	modeName := req.ModeName

	err := s.logic.ReportModePlaying(ctx, areaCode, modeName)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to report mode playing: %v", err)
	}


	return &pb.ModePlayingResponse{
		Status: "success",
	}, nil
}
package handlers

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	pb "multiplayer-modes-service/internal/models"
)

type MockBusinessLogic struct {
	mock.Mock
}

func (m *MockBusinessLogic) ReportModePlaying(ctx context.Context, areaCode, modeName string) error {
	args := m.Called(ctx, areaCode, modeName)
	return args.Error(0)
}

func (m *MockBusinessLogic) GetPopularModes(ctx context.Context, areaCode string) ([]*pb.Mode, error) {
	args := m.Called(ctx, areaCode)
	return args.Get(0).([]*pb.Mode), args.Error(1)
}

func TestGetPopularModesHandler(t *testing.T) {
	mockLogic := new(BusinessLogic)
	handler := &MultiplayerServiceServer{logic: mockLogic}

	mockModes := []*pb.Mode{
		{Name: "Battle Royale", Count: 10},
	}

	mockLogic.On("GetPopularModes", mock.Anything, "123").Return(mockModes, nil)

	req := &pb.GetPopularModesRequest{AreaCode: "123"}
	resp, err := handler.GetPopularModes(context.Background(), req)
	require.NoError(t, err, "GetPopularModes Failed")
	require.Equal(t, mockModes, resp.Modes)
}

func TestReportModePlayingHandler(t *testing.T) {
	mockLogic := new(MockBusinessLogic)
	handler := &MultiplayerServiceServer{logic: mockLogic}

	mockLogic.On("ReportModePlaying", mock.Anything, "123", "Battle Royale").Return(nil)

	req := &pb.ReportModePlayingRequest{AreaCode: "123", ModeName: "Battle Royale"}
	_, err := handler.ReportModePlaying(context.Background(), req)
	require.NoError(t, err, "ReportModePlaying handler failed")
	require.Equal(t, "success", resp.Status)
}
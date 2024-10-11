package business

import (
	"testing"

	"github.com/stretchr/testify/require"
	"multiplayer-modes-service/internal/test"
)

func TestReportModePlaying(t *testing.T) {
	test.LoadEnv(t)
	logic:= NewBusinessLogic()

	err := logic.ReportModePlaying(test.GetTestContext(),"123", "Battle Royale")
	require.NoError(t, err, "ReportModePlaying Failed")
}

func TestGetPopularModes(t *testing.T) {
	test.LoadEnv(t)
	logic:= NewBusinessLogic()

	_, err := logic.GetPopularModes(test.GetTestContext(),"123")
	require.NoError(t, err, "GetPopularModes Failed")
	require.NotNil(t, modes, "Expected modes, got nil")
}

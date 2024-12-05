package storage

import (
	"testing"

	"github.com/stretchr/testify/require"
	"multiplayer-modes-service/internal/test"
)

func TestIncrementModeCount(t *testing.T) {
	test.LoadEnv(t)
	storage := GetStorageInstance()
	err := storage.IncrementModeCount("123", "Battle Royale")
	require.NoError(t, err, "IncrementModeCount Failed")
}

func TestGetPopularModes(t *testing.T) {
	test.LoadEnv(t)
	storage := GetStorageInstance()
	_, err := storage.GetPopularModes("123")
	require.NoError(t, err, "GetPopularModes Failed")
}
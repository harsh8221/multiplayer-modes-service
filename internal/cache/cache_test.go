package cache 

import (
	"testing"

	"github.com/stretchr/testify/require"
	"multiplayer-modes-service/internal/test"
)

func TestSetAndGetPopularModes(t *testing.T) {
	test.LoadEnv(t)
	cache := GetCacheInstance()

	modes := []string{"Battle Royale", "Deathmatch"}
	err := cache.SetPopularModes("123", modes)
	require.NoError(t, err, "SetPopularModes Failed")

	data, err := cache.GetPopularModes("123")
	require.NoError(t, err, "GetPopularModes Failed")
	require.NotNil(t, data, "GetPopularModes returned nil")
}

func TestInvalidatePopularModes(t *testing.T) {
	test.LoadEnv(t)
	cache := GetCacheInstance()

	modes := []string{"Battle Royale", "Deathmatch"}
	err := cache.SetPopularModes("123", modes)
	require.NoError(t, err, "SetPopularModes Failed")

	err = cache.InvalidatePopularModes("123")
	require.NoError(t, err, "InvalidatePopularModes Failed")

	data, err := cache.GetPopularModes("123")
	require.NoError(t, err, "GetPopularModes Failed")
	require.Nil(t, data, "GetPopularModes returned non-nil data")
}
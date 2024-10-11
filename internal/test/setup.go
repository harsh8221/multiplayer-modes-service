package test

import (
	"context"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func LoadEnv(t *testing.T) {
	err := godotenv.Load("../../.env")
	require.NoError(t, err, "Failed to load Environment Variables")
}

func GetTestContext() context.Context {
	return context.Background()
}
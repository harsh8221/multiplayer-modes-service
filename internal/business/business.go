package business

import (
	"context"
	"encoding/json"
	"log"

	"multiplayer-modes-service/internal/cache"
	"multiplayer-modes-service/internal/storage"
	pb "multiplayer-modes-service/internal/models"
)

type BusinessLogic struct {
	cache *cache.Cache
	storage *storage.Storage
}

func NewBusinessLogic() *BusinessLogic {
	return &BusinessLogic{
		cache: cache.GetCacheInstance(),
		storage: storage.GetStorageInstance(),
	}
}

func (b *BusinessLogic) ReportModePlaying(ctx context.Context, areaCode, modeName string) error {

	err := b.storage.IncrementModeCount(areaCode, modeName)
	if err != nil {
		return err
	}

	err = b.cache.InvalidatePopularModes(areaCode)
	if err != nil {
		log.Printf("Failed to invalidate cache: %v", err)
	}
	return nil
}

func (b *BusinessLogic) GetPopularModes(ctx context.Context, areaCode string) ([]*pb.Mode, error) {

	data, err := b.cache.GetPopularModes(areaCode)
	if err != nil {
		log.Printf("Failed to get popular modes from cache: %v", err)
	}

	if data != nil {
		var modes []*pb.Mode
		err = json.Unmarshal(data, &modes)
		if err != nil {
			return modes, err
		}
		log.Printf("Cache hit for popular modes: %v", err)
	}

	// Cache miss or unmarshalling error, fetch from storage
	modeCounts, err := b.storage.GetPopularModes(areaCode)
	if err != nil {
		return nil, err
	}

	var modes []*pb.Mode
	for _, mc := range modeCounts {
		modes = append(modes, &pb.Mode{
			Name: mc.ModeName,
			PlayerCount: mc.PlayerCount,
		})
	}

	err = b.cache.SetPopularModes(areaCode, modes)
	if err != nil {
		log.Printf("Failed to set popular modes in cache: %v", err)
	}

	return modes, nil
}
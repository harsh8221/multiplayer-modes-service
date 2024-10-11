package cache

import (
	"context"
	"encoding/json"
	"os"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	client *redis.Client
}

var instance *Cache
var once sync.Once

// GetCacheInstance ensures a singleton instance of Cache
func GetCacheInstance() *Cache {
	once.Do(func () {
		redisAddr := os.Getenv("REDIS_ADDRESS")

		client := redis.NewClient(&redis.Options{
			Addr: redisAddr,
			Password: "",
			DB: 0,
		})

		instance = &Cache{
			client: client,
		}
	})
	return instance
}


func (c *Cache) SetPopularModes(areaCode string, modes interface{}) error {
	ctx := context.Background()
	key := "popular_modes" + areaCode

	data, err := json.Marshal(modes)
	if err != nil {
		return err
	}

	return c.client.Set(ctx, key, data, 5*time.Minute).Err()
}

func (c *Cache) GetPopularModes(areaCode string) ([]byte, error) {
	ctx := context.Background()
	key := "popular_modes" + areaCode

	data, err := c.client.Get(ctx, key).Bytes()
	if err == redis.Nil {
		// Cache miss
		return nil, nil
	} else if err != nil {
		return nil, err
	}


	return []byte(data), nil
}

func (c *Cache) InvalidatePopularModes(areaCode string) error {
	ctx := context.Background()
	key := "popular_modes" + areaCode

	return c.client.Del(ctx, key).Err()
}
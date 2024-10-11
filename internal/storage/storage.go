package storage

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

type Storage struct {
	client *mongo.Client
	collection *mongo.Collection
}

var instance *Storage
var once sync.Once

type ModeCount struct {
	AreaCode string `bson:"area_code"`
	ModeName string `bson:"mode_name"`
	PlayerCount int32 `bson:"player_count"`
}

// GetStorageInstance ensures a singleton instance of Storage
func GetStorageInstance() *Storage {
	once.Do(func() {
		mongoURI := os.Getenv("MONGODB_URI")
		mongoDBName := os.Getenv("MONGODB_DATABASE")

		clientOptions := options.Client().ApplyURI(mongoURI)
		client, err := mongo.NewClient(clientOptions)
		if err != nil {
			log.Fatalf("Failed to create MongoDB client: %v", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err = client.Connect(ctx)
		if err != nil {
			log.Fatalf("Failed to connect to MongoDB: %v", err)
		}

		collection := client.Database(mongoDBName).Collection("mode_counts")

		instance = &Storage{
			client: client,
			collection: collection,
		}

	})
	return instance
}

func (s *Storage) IncrementModeCount(areaCode, modeName string) error {
	filter := bson.M{
		"area_code": areaCode,
		"mode_name": modeName,
	}
	update := bson.M{
		"$inc": bson.M{"player_count": 1},
	}
	opts := options.Update().SetUpsert(true)
	_, err := s.collection.UpdateOne(context.Background(), filter, update, opts)
	return err
}

func (s *Storage) GetPopularModes(areaCode string) ([]ModeCount, error) {
	filter := bson.M{"area_code": areaCode}
	opts := options.Find().SetSort(bson.D{{"player_count", -1}}).SetLimit(10)

	cursor, err := s.collection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	var modes []ModeCount
	for cursor.Next(context.Background()) {
		var mode ModeCount 
		if err := cursor.Decode(&mode); err != nil {
			return nil, err
		}
		modes = append(modes, mode)
	}
	return modes, nil
}


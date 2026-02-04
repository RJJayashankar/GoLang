package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Crop struct {
	Name       string    `bson:"name" json:"name"`
	PricePerKg float64   `bson:"price_per_kg" json:"price_per_kg"`
	Category   string    `bson:"category" json:"category"`
	UpdatedAt  time.Time `bson:"updated_at" json:"updated_at"`
}

func main() {
	ctx := context.Background()

	// 1. MONGODB
	mClient, _ := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	mColl := mClient.Database("test").Collection("crops")

	// 2. CLOUD REDIS
	cloudURL := ""
	opt, _ := redis.ParseURL(cloudURL)
	rdb := redis.NewClient(opt)

	// 3. LOGIC
	cropName := "Rice"

	// Check Redis
	cachedPrice, err := rdb.Get(ctx, cropName).Result()
	if err == nil {
		fmt.Println("🚀 CACHE HIT (Redis): ₹" + cachedPrice)
		return
	}

	// Cache Miss
	fmt.Println("🔍 CACHE MISS: Searching MongoDB...")
	var crop Crop
	err = mColl.FindOne(ctx, bson.M{"name": cropName}).Decode(&crop)
	if err != nil {
		log.Fatal("Could not find or decode:", err)
	}

	// 4. UPDATE REDIS
	// Use the new field name here too
	priceStr := fmt.Sprintf("%.2f", crop.PricePerKg)
	rdb.Set(ctx, cropName, priceStr, 10*time.Minute)

	fmt.Printf("✅ DATABASE HIT: %s (%s) is ₹%s\n", crop.Name, crop.Category, priceStr)
}

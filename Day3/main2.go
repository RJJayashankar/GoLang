package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Crop struct {
	Name       string    `bson:"name"`
	Category   string    `bson:"category"`
	PricePerKg float64   `bson:"price_per_kg"`
	UpdatedAt  time.Time `bson:"updated_at"`
}

// Global collection handle for convenience in this script
var mColl *mongo.Collection

func main() {
	ctx := context.Background()

	// --- CONNECTION ---
	mClient, _ := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	mColl = mClient.Database("test").Collection("crops")

	// --- ADDING MULTIPLE CROPS ---
	// You can call this function for any new crop you want to add
	upsertCrop(ctx, "Potato", "Vegetable", 22.00)
	upsertCrop(ctx, "Drumstick", "vegetable", 55.50)
	upsertCrop(ctx, "cabage", "Vegetable", 70.00)
	upsertCrop(ctx, "bottle gourd", "Vegetable", 40.75)

	fmt.Println("\nAll crops have been synchronized with MongoDB!")
}

// upsertCrop will update the price if the crop exists, or create it if it doesn't.
func upsertCrop(ctx context.Context, name string, category string, price float64) {
	newCrop := Crop{
		Name:       name,
		Category:   category,
		PricePerKg: price,
		UpdatedAt:  time.Now(),
	}

	filter := bson.M{"name": name}
	update := bson.M{"$set": newCrop}
	opts := options.UpdateOne().SetUpsert(true)

	_, err := mColl.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Printf("Failed to add %s: %v", name, err)
	} else {
		fmt.Printf("Successfully added/updated: %s at ₹%.2f\n", name, price)
	}
}

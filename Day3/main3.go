package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// Crop struct matches your MongoDB Compass fields
type Crop struct {
	Name       string    `bson:"name" json:"name"`
	PricePerKg float64   `bson:"price_per_kg" json:"price_per_kg"`
	Category   string    `bson:"category" json:"category"`
	UpdatedAt  time.Time `bson:"updated_at" json:"updated_at"`
}

var mColl *mongo.Collection
var rdb *redis.Client

func main() {
	// 1. MONGODB CONNECTION
	// We use context.TODO() here since it's a one-time setup
	mClient, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("MongoDB Connection Error:", err)
	}
	mColl = mClient.Database("test").Collection("crops")

	// 2. CLOUD REDIS CONNECTION
	// Replace with your actual rediss:// connection string
	cloudURL := ""
	opt, err := redis.ParseURL(cloudURL)
	if err != nil {
		log.Fatal("Invalid Redis URL:", err)
	}
	rdb = redis.NewClient(opt)

	// 3. DEFINE THE ROUTE
	http.HandleFunc("/crop", getCropHandler)

	// 4. START SERVER
	fmt.Println("🚀 Server starting at http://localhost:8080")
	fmt.Println("Try: http://localhost:8080/crop?name=Ragi")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

func getCropHandler(w http.ResponseWriter, r *http.Request) {
	// Use the context from the incoming request
	ctx := r.Context()

	// Get the "name" parameter from the URL (?name=Ragi)
	cropName := r.URL.Query().Get("name")
	if cropName == "" {
		http.Error(w, "Please provide a crop name, e.g., /crop?name=Ragi", http.StatusBadRequest)
		return
	}

	// Step A: Check Redis (Cache)
	cachedPrice, err := rdb.Get(ctx, cropName).Result()
	if err == nil {
		fmt.Fprintf(w, "🚀 CACHE HIT (Cloud Redis)\nCrop: %s\nPrice: ₹%s", cropName, cachedPrice)
		return
	}

	// Step B: Cache Miss - Check MongoDB
	var crop Crop
	err = mColl.FindOne(ctx, bson.M{"name": cropName}).Decode(&crop)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Crop not found in MongoDB", http.StatusNotFound)
		} else {
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
		return
	}

	// Step C: Update Redis for next time (Expires in 10 minutes)
	priceStr := fmt.Sprintf("%.2f", crop.PricePerKg)
	rdb.Set(ctx, cropName, priceStr, 10*time.Minute)

	// Send Response to Postman/Browser
	fmt.Fprintf(w, "✅ DATABASE HIT (MongoDB)\nName: %s\nCategory: %s\nPrice: ₹%s",
		crop.Name, crop.Category, priceStr)
}

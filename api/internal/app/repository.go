package app

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	apimodel "github.com/gonza56d/gauth/pkg"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var uri = "mongodb://" + os.Getenv("MONGO_INITDB_ROOT_USERNAME") + ":" + os.Getenv("MONGO_INITDB_ROOT_PASSWORD") + "@" + os.Getenv("MONGO_INITDB_HOST") + ":" + os.Getenv("MONGO_INITDB_PORT")
var database = os.Getenv("MONGO_INITDB_DATABASE")
var collection = os.Getenv("MONGO_INITDB_COLLECTION")

// withMongoClient is a helper function that manages the lifecycle of a MongoDB client,
// providing a convenient way to execute operations on a MongoDB collection.
//
// Parameters:
// - callback: A function that accepts a *mongo.Collection and returns an error. This function
//   contains the actual database operations to be performed.
//
// Returns:
// - An error returned by the callback function, if any.
func withMongoClient(callback func(client *mongo.Collection) error) error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database(database).Collection(collection)
	return callback(coll)
}

func login(request *apimodel.LoginRequest) bool {
	var count int64
	err := withMongoClient(func(coll *mongo.Collection) error {
		var err error
		count, err = coll.CountDocuments(context.TODO(), bson.D{
			{Key: "email", Value: request.Email},
			{Key: "password", Value: request.Password},
		})
		if err != nil {
			panic(err)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	if count > 0 {
		return true
	}
	return false
}

func getRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
	})
	return rdb
}

func storeJWT(userEmail string, jwt string) {
	rdb := getRedisClient()
	tokenExpStr := os.Getenv("JWT_EXPIRATION_TIME_IN_HOURS")
	if tokenExpStr == "" {
		panic("JWT_EXPIRATION_TIME_IN_HOURS is not set")
	}
	tokenExp, err := strconv.Atoi(tokenExpStr)
	if err != nil {
		panic(err)
	}
	err = rdb.Set(nil, userEmail, jwt, time.Duration(tokenExp) * time.Hour).Err()
	if err!= nil {
		panic(err)
	}
}

func getJWT(userEmail string) string {
	rdb := getRedisClient()
	val, err := rdb.Get(nil, userEmail).Result()
	if err != nil {
		panic(err)
	}
	return val
}

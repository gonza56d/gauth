package app

import (
	"context"
	"os"

	//apimodel "github.com/gonza56d/gauth/pkg"
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

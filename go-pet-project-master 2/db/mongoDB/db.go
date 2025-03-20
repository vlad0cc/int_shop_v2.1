package mongoDB

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoProductCollection(host string) *mongo.Collection {
	client, err := mongo.NewClient(options.Client().ApplyURI(host))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("products").Collection("products")
	return collection
}

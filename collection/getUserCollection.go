package collection

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserCollection(client *mongo.Client, collectionUserName string) *mongo.Collection {
	userCollection := client.Database("TaskApp").Collection(collectionUserName)
	return userCollection
}

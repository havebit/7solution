package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoDB struct {
	client *mongo.Client
}

func (m *MongoDB) Save(originalURL, shortenURL string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	collection := m.client.Database("urlshortener").Collection("urls")
	_, err := collection.InsertOne(ctx, bson.M{"key": originalURL, "value": shortenURL})
	return err
}

func (m *MongoDB) OriginalURL(shortenURL string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	collection := m.client.Database("urlshortener").Collection("urls")

	var v result

	err := collection.FindOne(ctx, bson.M{"key": shortenURL}).Decode(&v)
	return v.Value, err
}

type result struct {
	Value string `bson:"value"`
}

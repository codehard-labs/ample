package mongo

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoURI          string        = ""
	MongoDBName       string        = "ample"
	EmptyFilter                     = &bson.M{}
	ConnectionTimeout time.Duration = 10 * time.Second
	MongoQueryTimeout time.Duration = 10 * time.Second
)

func GetMongoClient(uri string) (*mongo.Client, func()) {
	if uri == "" {
		uri = MongoURI
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil || client == nil {
		log.Print("mongo connect error", err)
		return nil, func() {}
	}
	return client, func() {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Println("mongo disconnect error")
		}
	}
}

// The input res should be a pointer
func FindOne(collectionName string, filter interface{}, res interface{}, opts ...*options.FindOneOptions) error {
	client, disconnect := GetMongoClient("")
	if client == nil {
		return errors.New("get mongo client failed")
	}
	defer disconnect()
	collection := client.Database(MongoDBName).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), MongoQueryTimeout)
	defer cancel()
	r := collection.FindOne(ctx, filter, opts...)
	return r.Decode(res)
}

// The input res should be a pointer
func FindAll(collectionName string, filter interface{}, res interface{}, opts ...*options.FindOptions) error {
	client, disconnect := GetMongoClient("")
	if client == nil {
		return errors.New("get mongo client failed")
	}
	defer disconnect()
	collection := client.Database(MongoDBName).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), MongoQueryTimeout)
	defer cancel()
	r, err := collection.Find(ctx, filter, opts...)
	if err != nil {
		return err
	}
	r.All(context.TODO(), res)
	return nil
}

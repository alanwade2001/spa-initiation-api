package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"k8s.io/klog/v2"
)

// MongoService s
type MongoService struct {
}

// NewMongoService s
func NewMongoService() RepositoryAPI {
	return &MongoService{}
}

// MongoConnection s
type MongoConnection struct {
	client *mongo.Client
	ctx    context.Context
	cancel context.CancelFunc
}

// Disconnect f
func (mc *MongoConnection) Disconnect() {
	mc.cancel()
	mc.client.Disconnect(mc.ctx)
}

// Connect f
func (ms MongoService) connect() MongoConnection {
	username := viper.GetString("MONGODB_USER")
	klog.Infof("mongo user: [%s]", username)
	password := viper.GetString("MONGODB_PASSWORD")
	uriTemplate := viper.GetString("MONGODB_URI_TEMPLATE")
	klog.Infof("uriTemplate: [%s]", uriTemplate)

	connectionURI := fmt.Sprintf(uriTemplate, username, password)

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		klog.Warningf("Failed to create client: %v", err)
	}

	connectTimeout := viper.GetDuration("MONGODB_TIMEOUT") * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)

	err = client.Connect(ctx)
	if err != nil {
		klog.Warningf("Failed to connect to cluster: %v", err)
	}

	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		klog.Warningf("Failed to ping cluster: %v", err)
	}

	klog.Infof("Connected to MongoDB!")

	return MongoConnection{client, ctx, cancel}
}

func (ms MongoService) getCollection(connection MongoConnection) *mongo.Collection {
	database := viper.GetString("MONGODB_DATABASE")
	return connection.client.Database(database).Collection("Initiations")
}

// CreateInitiation f
func (ms MongoService) CreateInitiation(initiation *Initiation) (*Initiation, error) {
	connection := ms.connect()
	defer connection.Disconnect()

	initiation.ID = primitive.NewObjectID().Hex()

	result, err := ms.getCollection(connection).InsertOne(connection.ctx, initiation)

	if err != nil {
		klog.Warningf("Could not create Initiation: %v", err)
		return nil, err
	}

	klog.Infof("result:[%+v]", result)

	return initiation, nil
}

// GetInitiation f
func (ms MongoService) GetInitiation(ID string) (*Initiation, error) {
	connection := ms.connect()
	defer connection.Disconnect()

	initiation := new(Initiation)
	filter := bson.M{"_id": ID}

	if err := ms.getCollection(connection).FindOne(connection.ctx, filter).Decode(initiation); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}

		return nil, err
	}

	klog.Infof("initiation:[%+v]", initiation)

	return initiation, nil
}

// GetInitiations f
func (ms MongoService) GetInitiations() (*Initiations, error) {
	connection := ms.connect()
	defer connection.Disconnect()

	var cursor *mongo.Cursor
	var err error
	var initiations Initiations

	filter := bson.M{}
	if cursor, err = ms.getCollection(connection).Find(connection.ctx, filter); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return &initiations, nil
		}

		return nil, err
	}

	if err = cursor.All(connection.ctx, &initiations); err != nil {
		return nil, err
	}

	return &initiations, nil
}

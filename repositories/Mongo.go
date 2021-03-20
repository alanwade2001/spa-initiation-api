package repositories

import (
	"errors"
	"reflect"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	mgo "github.com/alanwade2001/spa-common/mongo"
	"github.com/alanwade2001/spa-initiation-api/generated/initiation"
	"github.com/alanwade2001/spa-initiation-api/types"

	"k8s.io/klog/v2"
)

// MongoRepository s
type MongoRepository struct {
	service *mgo.MongoService
}

// NewMongoRepository s
func NewMongoRepository() types.RepositoryAPI {
	return &MongoRepository{}
}

func (ms *MongoRepository) GetService() *mgo.MongoService {

	if ms.service != nil {
		return ms.service
	}

	uriTemplate := viper.GetString("MONGODB_URI_TEMPLATE")
	username := viper.GetString("MONGODB_USER")
	password := viper.GetString("MONGODB_PASSWORD")
	connectTimeout := viper.GetDuration("MONGODB_TIMEOUT") * time.Second
	database := viper.GetString("MONGODB_DATABASE")
	collection := viper.GetString("MONGODB_COLLECTION")

	structcodec, _ := bsoncodec.NewStructCodec(bsoncodec.JSONFallbackStructTagParser)
	reg := bson.NewRegistryBuilder().
		RegisterTypeEncoder(reflect.TypeOf(initiation.InitiationModel{}), structcodec).
		RegisterTypeDecoder(reflect.TypeOf(initiation.InitiationModel{}), structcodec).
		Build()

	service := mgo.NewMongoService(uriTemplate, username, password, database, collection, connectTimeout, reg)

	ms.service = service

	return ms.service
}

// CreateInitiation f
func (ms MongoRepository) CreateInitiation(init *initiation.InitiationModel) (*initiation.InitiationModel, error) {
	connection := ms.GetService().Connect()
	defer connection.Disconnect()

	init.Id = primitive.NewObjectID().Hex()

	result, err := ms.GetService().GetCollection(connection).InsertOne(connection.Ctx, init)

	if err != nil {
		klog.Warningf("Could not create Initiation: %v", err)
		return nil, err
	}

	klog.Infof("result:[%+v]", result)

	return init, nil
}

// GetInitiation f
func (ms MongoRepository) GetInitiation(ID string) (*initiation.InitiationModel, error) {
	connection := ms.GetService().Connect()
	defer connection.Disconnect()

	init := new(initiation.InitiationModel)
	filter := bson.M{"_id": ID}

	if err := ms.GetService().GetCollection(connection).FindOne(connection.Ctx, filter).Decode(init); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}

		return nil, err
	}

	klog.Infof("initiation:[%+v]", init)

	return init, nil
}

// GetInitiations f
func (ms MongoRepository) GetInitiations() ([]*initiation.InitiationModel, error) {
	connection := ms.GetService().Connect()
	defer connection.Disconnect()

	var cursor *mongo.Cursor
	var err error
	initiations := []*initiation.InitiationModel{}

	filter := bson.M{}
	if cursor, err = ms.GetService().GetCollection(connection).Find(connection.Ctx, filter); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return initiations, nil
		}

		return nil, err
	}

	defer cursor.Close(connection.Ctx)

	if err = cursor.All(connection.Ctx, &initiations); err != nil {
		return nil, err
	}

	return initiations, nil
}

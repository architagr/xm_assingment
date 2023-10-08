package persistance

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ICollection[T any] interface {
	GetById(id string) (data *T, err error)
	Get(filter primitive.M, pageSize int64, startPage int64) (data []T, err error)
	AddSingle(data T) (id interface{}, err error)
	Delete(id string) (err error)
	UpdateSingle(data T, _id string) error
}

type collection[T any] struct {
	client                       IConnection
	databaseName, collectionName string
}

func InitCollection[T any](conn IConnection, databaseName, collectionName string) (ICollection[T], error) {

	return &collection[T]{
		client:         conn,
		databaseName:   databaseName,
		collectionName: collectionName,
	}, nil
}

func (doc *collection[T]) getCollection() (collection *mongo.Collection, err error) {
	client, _, err := doc.client.GetConnction()
	if err != nil {
		return
	}
	collection = client.Database(doc.databaseName).Collection(doc.collectionName)
	return
}
func (doc *collection[T]) Get(filter primitive.M, pageSize int64, startPage int64) (data []T, err error) {
	if pageSize == 0 {
		pageSize = 10
	}
	skip := startPage * pageSize
	if skip > 0 {
		skip--
	}
	filterOptions := options.Find()
	filterOptions.Limit = &pageSize
	filterOptions.Skip = &skip
	collection, err := doc.getCollection()
	if err != nil {
		return
	}

	result, err := collection.Find(context.TODO(), filter, filterOptions)
	if err != nil {
		return
	}
	data = make([]T, 0)
	err = result.All(context.TODO(), &data)
	if err != nil {
		return
	}
	return
}
func (doc *collection[T]) GetById(id string) (data *T, err error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}

	collection, err := doc.getCollection()
	if err != nil {
		return
	}
	data = new(T)
	err = collection.FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&data)
	return
}

func (doc *collection[T]) AddSingle(data T) (id interface{}, err error) {
	collection, err := doc.getCollection()
	if err != nil {
		return
	}
	result, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		return
	}
	id = result.InsertedID
	return
}

func (doc *collection[T]) Delete(id string) (err error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}

	collection, err := doc.getCollection()
	if err != nil {
		return
	}

	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": objectId})
	if err != nil {
		return
	}

	return
}

func (doc *collection[T]) UpdateSingle(data T, _id string) error {
	objectId, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		return err
	}
	collection, err := doc.getCollection()
	if err != nil {
		return err
	}

	_, err = collection.UpdateOne(context.Background(), bson.M{"_id": objectId}, bson.M{"$set": data})
	if err != nil {
		return err
	}
	return nil
}

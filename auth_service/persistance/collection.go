package persistance

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type ICollection[T any] interface {
	Get(filter bson.M) (data *T, err error)
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

func (doc *collection[T]) Get(filter bson.M) (data *T, err error) {
	client, _, err := doc.client.GetConnction()
	if err != nil {
		return
	}
	collection := client.Database(doc.databaseName).Collection(doc.collectionName)
	data = new(T)
	err = collection.FindOne(context.Background(), filter).Decode(&data)
	return
}

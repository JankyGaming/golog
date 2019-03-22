package golog

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

//NewLogClientMongo returns a client for writing logs and errors to a mongo collection
func NewLogClientMongo(client *mongo.Client, dbName string, collectionName string) (LogClient, error) {
	logCollection := client.Database(dbName).Collection(collectionName)

	return &mongoLogClient{MongoCli: client, LogCollection: logCollection}, nil
}

//Log writes a log to the database
func (c *mongoLogClient) Log(t logType, l logLevel, desc string, data map[string]interface{}) {
	newLog := buildLog(t, l, desc, data)

	_, err := c.LogCollection.InsertOne(context.Background(), newLog)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(newLog.StdOutPrint)
	if data != nil {
		fmt.Printf("%+v\n", data)
	}
}

type mongoLogClient struct {
	MongoCli      *mongo.Client
	LogCollection *mongo.Collection
}

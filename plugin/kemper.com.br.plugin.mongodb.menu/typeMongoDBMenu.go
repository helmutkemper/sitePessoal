package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

var Menu = MongoDBMenu{}

type MongoDBMenu struct {
	Client     *mongo.Client
	Ctx        context.Context
	CancalFunc context.CancelFunc
	ClientMenu *mongo.Collection
}

package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

var Language MongoDBLanguage

type MongoDBLanguage struct {
	Client         *mongo.Client
	Ctx            context.Context
	CancalFunc     context.CancelFunc
	ClientLanguage *mongo.Collection
}

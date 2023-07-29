package mongodb

import (
	"context"
	"log"

	"github.com/amookia/arvan-backend-challenge/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepo struct {
	db      *mongo.Database
	logger  *log.Logger
	context context.Context
}

func New(conn *mongo.Database, logger *log.Logger, ctx context.Context) repository.Mongodb {
	conn.Collection(objectsName).Indexes().CreateMany(
		context.Background(),
		[]mongo.IndexModel{
			{Keys: bson.D{{Key: "checksum", Value: 1}},
				Options: options.Index().SetUnique(true)},
			{Keys: bson.D{{Key: "uuid", Value: 1}},
				Options: options.Index().SetUnique(true)},
		},
	)
	return mongoRepo{db: conn, logger: logger, context: ctx}
}

func (m mongoRepo) IsDupKey(err error, key string) bool {
	doc := err.(mongo.WriteException).WriteErrors[0].Raw.Lookup("keyPattern").
		Document().Lookup(key).String()
	return doc != ""
}

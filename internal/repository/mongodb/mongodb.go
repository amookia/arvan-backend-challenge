package mongodb

import (
	"context"
	"log"

	"github.com/amookia/arvan-backend-challenge/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepo struct {
	db      *mongo.Database
	logger  *log.Logger
	context context.Context
}

func New(conn *mongo.Database, logger *log.Logger, ctx context.Context) repository.Mongodb {
	return mongoRepo{db: conn, logger: logger, context: ctx}
}

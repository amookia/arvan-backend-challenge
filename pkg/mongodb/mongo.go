package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectToMongo func connects to mongo db
func ConnectToMongo(mongoURL string, DbName string) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	db := client.Database(DbName)
	return db, nil
}

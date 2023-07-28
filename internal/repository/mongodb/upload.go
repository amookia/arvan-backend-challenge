package mongodb

import (
	"github.com/amookia/arvan-backend-challenge/internal/entity/object"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const collectionName = "objects"

func (m mongoRepo) InsertObject(data object.ObjectModel) (primitive.ObjectID, error) {
	collection := m.db.Collection(collectionName)
	res, err := collection.InsertOne(m.context, data)
	if err != nil {
		m.logger.Println(err)
		return primitive.NilObjectID, err
	}
	id := res.InsertedID.(primitive.ObjectID)
	return id, nil
}

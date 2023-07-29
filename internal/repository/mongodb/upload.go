package mongodb

import (
	"errors"

	"github.com/amookia/arvan-backend-challenge/internal/entity/object"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const objectsName = "objects"

func (m mongoRepo) InsertObject(data object.ObjectModel) (primitive.ObjectID, error) {
	collection := m.db.Collection(objectsName)
	res, err := collection.InsertOne(m.context, data)
	if err != nil {
		switch m.IsDupKey(err, "uuid") {
		case true:
			return primitive.NilObjectID, errors.New("duplicate uuid")
		case false:
			return primitive.NilObjectID, errors.New("duplicate object")
		}
		return primitive.NilObjectID, err
	}
	id := res.InsertedID.(primitive.ObjectID)
	return id, nil
}

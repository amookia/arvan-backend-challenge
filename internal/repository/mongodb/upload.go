package mongodb

import (
	"errors"
	"fmt"

	"github.com/amookia/arvan-backend-challenge/internal/entity/object"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
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

func (m mongoRepo) DeleteObjectByObjectIdUser(username string, objectId uuid.UUID) (bool, error) {
	collection := m.db.Collection(objectsName)
	fmt.Println(username, objectId)
	res, err := collection.DeleteOne(m.context, bson.M{"owner": username, "uuid": objectId})
	if err != nil {
		return false, err
	}
	return res.DeletedCount == 1, nil
}

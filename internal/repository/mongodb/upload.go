package mongodb

import (
	"errors"

	"github.com/amookia/arvan-backend-challenge/internal/entity/object"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const objectsName = "objects"
func (m mongoRepo) IsChecksumExists(data object.ObjectModel) (bool,error) {
	collection := m.db.Collection(objectsName)
	count,err := collection.CountDocuments(
		m.context,bson.M{"checksum":data.CheckSum,"size":data.Size},
	)
	return count > 0,err

}

func (m mongoRepo) InsertObject(data object.ObjectModel) (primitive.ObjectID, error) {
	collection := m.db.Collection(objectsName)
	res, err := collection.InsertOne(m.context, data)
	if err != nil {
		switch {
		case m.IsDupKey(err, "uuid") :
			return primitive.NilObjectID, errors.New("duplicate uuid")
		}
		return primitive.NilObjectID, err
	}
	id := res.InsertedID.(primitive.ObjectID)
	return id, nil
}

func (m mongoRepo) DeleteObjectByObjectIdUser(username string, objectId uuid.UUID) (bool, error) {
	collection := m.db.Collection(objectsName)
	res, err := collection.DeleteOne(m.context, bson.M{"owner": username, "uuid": objectId})
	if err != nil {
		return false, err
	}
	return res.DeletedCount == 1, nil
}
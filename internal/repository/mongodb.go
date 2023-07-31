package repository

import (
	"github.com/amookia/arvan-backend-challenge/internal/entity/object"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mongodb interface {
	InsertObject(object.ObjectModel) (primitive.ObjectID, error)
	DeleteObjectByObjectIdUser(string, uuid.UUID) (bool, error)
}

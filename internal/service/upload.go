package service

import (
	"github.com/amookia/arvan-backend-challenge/internal/transport/http/request"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Upload interface {
	CreateObject(request.PutObject) (primitive.ObjectID, error)
}

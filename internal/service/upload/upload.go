package upload

import (
	"github.com/amookia/arvan-backend-challenge/internal/entity/object"
	"github.com/amookia/arvan-backend-challenge/internal/repository"
	"github.com/amookia/arvan-backend-challenge/internal/service"
	"github.com/amookia/arvan-backend-challenge/internal/transport/http/request"
	"github.com/amookia/arvan-backend-challenge/pkg/checksum"
	"github.com/amookia/arvan-backend-challenge/pkg/logger"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type upload struct {
	mongodb repository.Mongodb
	logger  logger.Logger
}

func New(logger logger.Logger, mongodb repository.Mongodb) service.Upload {
	return upload{logger: logger, mongodb: mongodb}
}

func (u upload) PutObject(req request.PutObject) (primitive.ObjectID, error) {
	id, err := uuid.Parse(req.Data.Id)
	if err != nil {
		return primitive.NilObjectID, err
	}
	file, err := req.File.Open()
	if err != nil {
		return primitive.NilObjectID, err
	}
	object := object.ObjectModel{
		CheckSum: checksum.GenerateMd5CheckSum(file),
		Uuid:     id,
		Size:     req.File.Size,
		Owner:    req.Data.Username,
	}

	objectId, err := u.mongodb.InsertObject(object)

	if err != nil {
		return primitive.NilObjectID, err
	}
	return objectId, nil
}

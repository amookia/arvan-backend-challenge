package upload

import (
	"errors"

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
	redis   repository.Redis
	logger  logger.Logger
}

func New(logger logger.Logger, mongodb repository.Mongodb, redis repository.Redis) service.Upload {
	return upload{logger: logger, redis: redis, mongodb: mongodb}
}

func (u upload) CreateObject(req request.PutObject) (primitive.ObjectID, error) {
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
		Owner:    req.Username,
	}

	objectId, err := u.mongodb.InsertObject(object)

	if err != nil {
		return primitive.NilObjectID, err
	}
	u.redis.UserMonthlyUsageUpdate(object.Owner, object.Size)
	return objectId, nil
}

func (u upload) DeleteObject(username string, objectId string) error {
	id, err := uuid.Parse(objectId)
	if err != nil {
		return err
	}
	deleted, err := u.mongodb.DeleteObjectByObjectIdUser(username, id)
	if err != nil {
		return err
	}
	if !deleted {
		return errors.New("object does not exist")
	}
	return nil
}

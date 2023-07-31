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
)

type upload struct {
	mongodb repository.Mongodb
	redis   repository.Redis
	logger  logger.Logger
}

func New(logger logger.Logger, mongodb repository.Mongodb, redis repository.Redis) service.Upload {
	return upload{logger: logger, redis: redis, mongodb: mongodb}
}

func (u upload) CreateObject(req request.PutObject) (string, error) {
	id, err := uuid.Parse(req.ObjectId)
	if err != nil {
		return "", err
	}
	file, err := req.File.Open()
	if err != nil {
		return "", err
	}
	object := object.ObjectModel{
		CheckSum: checksum.GenerateMd5CheckSum(file),
		Uuid:     id,
		Size:     req.File.Size,
		Owner:    req.Username,
	}

	_, err = u.mongodb.InsertObject(object)

	if err != nil {
		return "", err
	}
	u.redis.UserMonthlyUsageUpdate(object.Owner, object.Size)
	return req.ObjectId, nil
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

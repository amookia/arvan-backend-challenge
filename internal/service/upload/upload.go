package upload

import (
	"log"

	"github.com/amookia/arvan-backend-challenge/internal/entity/object"
	"github.com/amookia/arvan-backend-challenge/internal/repository"
	"github.com/amookia/arvan-backend-challenge/internal/service"
	"github.com/amookia/arvan-backend-challenge/internal/transport/http/request"
	"github.com/amookia/arvan-backend-challenge/pkg/checksum"
	"github.com/google/uuid"
)

type upload struct {
	mongodb repository.Mongodb
	logger  *log.Logger
}

func New(logger *log.Logger, mongodb repository.Mongodb) service.Upload {
	return upload{logger: logger, mongodb: mongodb}
}

func (u upload) PutObject(req request.PutObject) error {
	id, err := uuid.Parse(req.Data.Id)
	if err != nil {
		return err
	}
	file, err := req.File.Open()
	if err != nil {
		return err
	}
	object := object.ObjectModel{
		CheckSum: checksum.GenerateMd5CheckSum(file),
		Uuid:     id,
		Size:     req.File.Size,
		Owner:    req.Data.Username,
	}

	_, err = u.mongodb.InsertObject(object)
	if err != nil {
		return err
	}
	return nil
}

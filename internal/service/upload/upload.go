package upload

import (
	"log"

	"github.com/amookia/arvan-backend-challenge/internal/repository"
	"github.com/amookia/arvan-backend-challenge/internal/service"
)

type upload struct {
	mongodb repository.Mongodb
	logger  *log.Logger
}

func New(logger *log.Logger, mongodb repository.Mongodb) service.Upload {
	return upload{logger: logger, mongodb: mongodb}
}

package service

import (
	"github.com/amookia/arvan-backend-challenge/internal/transport/http/request"
)

type Upload interface {
	CreateObject(request.PutObject) (string, error)
	DeleteObject(string, string) error
}

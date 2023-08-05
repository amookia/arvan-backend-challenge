package service

import (
	"github.com/amookia/arvan-backend-challenge/internal/transport/http/request"
)

type Upload interface {
	CreateObject(request.CreateObject) (string, error)
	DeleteObject(string, string) error
	PutObject(request.PutObject) error
}
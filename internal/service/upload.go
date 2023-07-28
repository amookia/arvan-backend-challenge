package service

import "github.com/amookia/arvan-backend-challenge/internal/transport/http/request"

type Upload interface {
	PutObject(request.PutObject) error
}

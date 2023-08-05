package request

import (
	"mime/multipart"
)

type (
	CreateObject struct {
		File     *multipart.FileHeader `form:"file" binding:"required"`
		ObjectId string                `form:"objectId"`
		Username string
	}
	PutObject struct {
		Body     []byte
		ObjectId string
		Username string
	}
)
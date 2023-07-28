package request

import "mime/multipart"

type (
	PutObject struct {
		File *multipart.FileHeader `form:"file" binding:"required"`
		Data Data                  `form:"data" binding:"required"`
	}
	Data struct {
		Id       string `json:"id" binding:"required"`
		Username string `json:"username"`
	}
)

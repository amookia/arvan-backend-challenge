package request

import "mime/multipart"

type (
	PutObject struct {
		File     *multipart.FileHeader `form:"file" binding:"required"`
		ObjectId string                `form:"objectId"`
		Username string
	}
)

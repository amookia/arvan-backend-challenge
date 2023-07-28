package request

import "mime/multipart"

type (
	RequestLimiter struct {
		Data struct {
			Username string `json:"username" binding:"required"`
		} `form:"data" binding:"required"`
	}

	MonthlyLimiter struct {
		File *multipart.FileHeader `form:"file" binding:"required"`
		Data struct {
			Username string `json:"username" binding:"required"`
		} `form:"data" binding:"required"`
	}
)

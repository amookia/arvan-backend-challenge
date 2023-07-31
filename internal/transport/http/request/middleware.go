package request

import "mime/multipart"

type (
	RequestLimiter struct {
		Username string `header:"username"`
		// Data     struct {
		// 	Username string `json:"username" binding:"required"`
		// } `form:"data" binding:"required"`
	}

	MonthlyLimiter struct {
		File     *multipart.FileHeader `form:"file" binding:"required"`
		Username string

		// Data struct {
		// 	Username string `json:"username" binding:"required"`
		// } `form:"data" binding:"required"`
	}
)

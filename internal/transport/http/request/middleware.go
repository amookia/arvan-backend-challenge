package request

type (
	RequestLimiter struct {
		Data struct {
			Username string `json:"username" binding:"required"`
		} `form:"data" binding:"required"`
	}
)

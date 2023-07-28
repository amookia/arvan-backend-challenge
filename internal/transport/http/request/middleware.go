package request

type (
	RequestLimiter struct {
		Username string `form:"username" binding:"required"`
	}
)

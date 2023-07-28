package gin

import (
	"log"

	"github.com/amookia/arvan-backend-challenge/internal/config"
	"github.com/amookia/arvan-backend-challenge/internal/service"
	"github.com/gin-gonic/gin"
)

type middlewareHandler struct {
	config     config.Middleware
	middleware service.Middleware
	logger     *log.Logger
}

func (m middlewareHandler) RequestLimiter(c *gin.Context) {
	user, _ := c.GetPostForm(m.config.UserKey)
	limited := m.middleware.IsUserLimited(user)
	if limited {
		c.AbortWithStatusJSON(429, gin.H{"error": "request limited"})
	}
	c.Next()

}

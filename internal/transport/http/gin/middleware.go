package gin

import (
	"log"

	"github.com/amookia/arvan-backend-challenge/internal/config"
	"github.com/amookia/arvan-backend-challenge/internal/service"
	"github.com/amookia/arvan-backend-challenge/internal/transport/http/request"
	"github.com/amookia/arvan-backend-challenge/internal/transport/http/response"
	"github.com/gin-gonic/gin"
)

type middlewareHandler struct {
	config     config.Middleware
	middleware service.Middleware
	logger     *log.Logger
}

func (m middlewareHandler) requestLimiter(c *gin.Context) {
	var form request.RequestLimiter
	err := c.ShouldBind(&form)
	if err != nil {
		c.AbortWithStatusJSON(400, response.RequestLimiter{Error: "invalid form"})
	}
	limited := m.middleware.IsUserLimited(form.Data.Username)
	if limited {
		c.AbortWithStatusJSON(429, response.RequestLimiter{Error: "request limited"})
	}
	c.Next()

}

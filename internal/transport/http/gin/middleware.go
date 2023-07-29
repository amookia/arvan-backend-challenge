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

func (m middlewareHandler) requestQuota(c *gin.Context) {
	var form request.RequestLimiter
	err := c.ShouldBind(&form)
	if err != nil {
		c.AbortWithStatusJSON(400, response.RequestLimiter{Error: "invalid form"})
	}
	limited := m.middleware.UserQuotaRequest(form.Data.Username)
	if limited {
		c.AbortWithStatusJSON(429, response.RequestLimiter{Error: "request limited"})
	}
	c.Next()

}

func (m middlewareHandler) monthlyQuota(c *gin.Context) {
	var form request.MonthlyLimiter
	err := c.ShouldBind(&form)
	if err != nil {
		c.AbortWithStatusJSON(400, response.RequestLimiter{Error: "invalid form"})
	}
	limited := m.middleware.UserQuotaTraffic(form.Data.Username, form.File.Size)
	if limited {
		c.AbortWithStatusJSON(429, response.RequestLimiter{Error: "monthly limited"})
	}
	c.Next()
}

package gin

import (
	"github.com/amookia/arvan-backend-challenge/internal/config"
	"github.com/amookia/arvan-backend-challenge/internal/service"
	"github.com/amookia/arvan-backend-challenge/internal/transport/http/request"
	"github.com/amookia/arvan-backend-challenge/internal/transport/http/response"
	"github.com/amookia/arvan-backend-challenge/pkg/logger"
	"github.com/gin-gonic/gin"
)

type middlewareHandler struct {
	config     config.Middleware
	middleware service.Middleware
	logger     logger.Logger
}

func (m middlewareHandler) requestQuota(c *gin.Context) {
	var form request.RequestLimiter
	err := c.ShouldBindHeader(&form)
	m.logger.Info(form.Username)
	if err != nil {
		m.logger.Error(err.Error())
		c.AbortWithStatusJSON(400, response.RequestLimiter{Error: "invalid form1"})
	}
	limited := m.middleware.UserQuotaRequest(form.Username)
	if limited {
		c.AbortWithStatusJSON(429, response.RequestLimiter{Error: "request limited"})
	}
	c.Next()

}

func (m middlewareHandler) monthlyQuota(c *gin.Context) {
	var form request.MonthlyLimiter
	form.Username = c.GetHeader("username")
	if form.Username == "" {
		c.AbortWithStatusJSON(400, response.RequestLimiter{Error: "invalid form"})
		return
	}
	limited := m.middleware.UserQuotaTraffic(form.Username)
	if limited {
		c.AbortWithStatusJSON(429, response.RequestLimiter{Error: "monthly limited"})
	}
	c.Next()
}

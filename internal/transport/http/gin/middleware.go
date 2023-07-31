package gin

import (
	"github.com/amookia/arvan-backend-challenge/internal/config"
	"github.com/amookia/arvan-backend-challenge/internal/service"
	"github.com/amookia/arvan-backend-challenge/internal/transport/http/request"
	"github.com/amookia/arvan-backend-challenge/internal/transport/http/response"
	"github.com/amookia/arvan-backend-challenge/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
	err := c.ShouldBindWith(&form, binding.FormMultipart)
	form.Username = c.GetHeader("username")
	if err != nil {
		m.logger.Error(err.Error())
		c.AbortWithStatusJSON(400, response.RequestLimiter{Error: "invalid form"})
	}
	limited := m.middleware.UserQuotaTraffic(form.Username, form.File.Size)
	if limited {
		c.AbortWithStatusJSON(429, response.RequestLimiter{Error: "monthly limited"})
	}
	c.Next()
}

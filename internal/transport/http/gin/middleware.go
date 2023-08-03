package gin

import (
	"errors"

	"github.com/amookia/arvan-backend-challenge/internal/config"
	"github.com/amookia/arvan-backend-challenge/internal/service"
	"github.com/amookia/arvan-backend-challenge/internal/transport/http/request"
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
		c.AbortWithStatusJSON(400, FormErrorHandler(err))
	}
	limited := m.middleware.UserQuotaRequest(form.Username)
	if limited {
		c.AbortWithStatusJSON(429, ErrorHandler("request limited",errors.New("too many requests in a minute")))
	}
	c.Next()

}

func (m middlewareHandler) monthlyQuota(c *gin.Context) {
	var form request.MonthlyLimiter
	form.Username = c.GetHeader("username")
	if form.Username == "" {
		c.AbortWithStatusJSON(400, ErrorHandler("invalid header",errors.New("username header required")))
		return
	}
	limited := m.middleware.UserQuotaTraffic(form.Username)
	if limited {
		c.AbortWithStatusJSON(429, ErrorHandler("user quota",errors.New("monthly limited quota")))
	}
	c.Next()
}
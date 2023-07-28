package gin

import (
	"log"

	"github.com/amookia/arvan-backend-challenge/internal/config"
	"github.com/amookia/arvan-backend-challenge/internal/service"
	"github.com/amookia/arvan-backend-challenge/internal/transport/http"
	"github.com/gin-gonic/gin"
)

type api struct {
	gin        *gin.Engine
	middleware middlewareHandler
}

func New(logger *log.Logger, mdConfig config.Middleware, mdService service.Middleware) http.Api {
	return &api{
		gin: gin.New(),
		middleware: middlewareHandler{
			config:     mdConfig,
			middleware: mdService,
			logger:     logger,
		},
	}
}

func (a api) Start(address string) error {
	routerGroup := a.gin.Group("/api")
	a.routes(routerGroup)
	return a.gin.Run(address)
}

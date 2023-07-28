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
	upload     uploadHandler
}

func New(logger *log.Logger, mdConfig config.Middleware,
	mdService service.Middleware, upService service.Upload) http.Api {
	return &api{
		gin: gin.New(),
		middleware: middlewareHandler{
			config:     mdConfig,
			middleware: mdService,
			logger:     logger,
		},
		upload: uploadHandler{
			upload: upService,
			logger: logger,
		},
	}
}

func (a api) Start(address string) error {
	v1RouterGroup := a.gin.Group("/api/v1")
	a.routes(v1RouterGroup)
	return a.gin.Run(address)
}

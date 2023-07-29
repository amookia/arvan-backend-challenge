package gin

import (
	"github.com/amookia/arvan-backend-challenge/internal/service"
	"github.com/amookia/arvan-backend-challenge/internal/transport/http/request"
	"github.com/amookia/arvan-backend-challenge/internal/transport/http/response"
	"github.com/amookia/arvan-backend-challenge/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type uploadHandler struct {
	upload service.Upload
	logger logger.Logger
}

func (u uploadHandler) Put(c *gin.Context) {
	var model request.PutObject
	c.BindWith(&model, binding.FormMultipart)
	err := u.upload.PutObject(model)
	if err != nil {
		c.AbortWithStatusJSON(400, response.PutObjectError{Err: err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "OK"})
}

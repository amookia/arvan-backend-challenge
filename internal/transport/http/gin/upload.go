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
	var form request.PutObject
	c.BindWith(&form, binding.FormMultipart)
	objectId, err := u.upload.PutObject(form)
	if err != nil {
		c.AbortWithStatusJSON(400, response.PutObjectError{Err: err.Error()})
		return
	}
	c.JSON(200, response.PutObject{Message: "success", ObjectId: objectId.Hex()})
}

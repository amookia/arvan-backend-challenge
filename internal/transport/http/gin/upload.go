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

func (u uploadHandler) Create(c *gin.Context) {
	var form request.PutObject
	err := c.ShouldBindWith(&form, binding.FormMultipart)
	form.Username = c.GetHeader("username")
	if err != nil {
		c.AbortWithStatusJSON(400, FormErrorHandler(err))
		return
	}
	objectId, err := u.upload.CreateObject(form)
	if err != nil {
		c.AbortWithStatusJSON(400, response.PutObjectError{Err: err.Error()})
		return
	}
	c.JSON(200, response.PutObject{Message: "success", ObjectId: objectId})
}

func (u uploadHandler) Delete(c *gin.Context) {
	objectId := c.Param("objectId")
	username := c.Request.Header.Get("username")
	err := u.upload.DeleteObject(username, objectId)
	if err != nil {
		c.AbortWithStatusJSON(404, response.Error{Errors: []string{err.Error()},
		Message: "failed to delete object"})
		return
	}
	c.JSON(200, response.PutObject{Message: "object has been deleted",ObjectId: objectId})
}

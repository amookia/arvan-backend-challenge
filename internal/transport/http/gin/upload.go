package gin

import (
	"errors"
	"io"
	"net/http"

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
	var form request.CreateObject
	err := c.ShouldBindWith(&form, binding.FormMultipart)
	form.Username = c.GetHeader("username")
	if err != nil {
		c.AbortWithStatusJSON(400, FormErrorHandler(err))
		return
	}
	objectId, err := u.upload.CreateObject(form)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusConflict,
			ErrorHandler("duplication error",err))
		return
	}
	c.JSON(
		http.StatusOK,
		response.PutObject{Message: "success", ObjectId: objectId},
	)
}

func (u uploadHandler) Delete(c *gin.Context) {
	objectId := c.Param("objectId")
	username := c.Request.Header.Get("username")
	err := u.upload.DeleteObject(username, objectId)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			response.Error{Errors: []string{err.Error()},Message: "failed to delete object"},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		response.PutObject{Message: "object has been deleted",ObjectId: objectId},
	)
}

func (u uploadHandler) Put(c *gin.Context) {
	requestBody := c.Request.Body
	if requestBody == nil {
		c.AbortWithStatusJSON(400,
		ErrorHandler("invalid input",errors.New("there is not any data input")))
	}
	body,err := io.ReadAll(requestBody)
	if err != nil {
		c.AbortWithStatusJSON(400,
		ErrorHandler("invalid inupt",errors.New("there is a problem to read input data")))
	}
	defer requestBody.Close()
	objectId := c.Param("objectId")
	username := c.Request.Header.Get("username")
	form := request.PutObject{
		Body: body,
		ObjectId: objectId,
		Username: username,
	}
	err = u.upload.PutObject(form)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusConflict, 
			ErrorHandler("duplication error",err),
		)
		return
	}
	c.JSON(
		http.StatusOK, 
		response.PutObject{Message: "success", ObjectId: objectId},
	)
}
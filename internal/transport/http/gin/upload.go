package gin

import (
	"fmt"
	"log"

	"github.com/amookia/arvan-backend-challenge/internal/service"
	"github.com/amookia/arvan-backend-challenge/internal/transport/http/request"
	"github.com/amookia/arvan-backend-challenge/internal/transport/http/response"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type uploadHandler struct {
	upload service.Upload
	logger *log.Logger
}

func (u uploadHandler) Put(c *gin.Context) {
	var model request.PutObject
	c.BindWith(&model, binding.FormMultipart)
	fmt.Println(model)
	err := u.upload.PutObject(model)
	if err != nil {
		c.AbortWithStatusJSON(400, response.PutObjectError{Err: err})
	}
	c.JSON(200, gin.H{"message": "OK"})
}

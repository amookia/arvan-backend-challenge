package gin

import (
	"log"

	"github.com/gin-gonic/gin"
)

type uploadHandler struct {
	logger *log.Logger
}

func (u uploadHandler) Put(c *gin.Context) {
	c.JSON(200, gin.H{"hello": "WORLD"})
}

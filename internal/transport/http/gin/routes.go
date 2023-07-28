package gin

import "github.com/gin-gonic/gin"

func (a api) routes(c *gin.RouterGroup) {
	c.POST("/api/hello", a.middleware.RequestLimiter, func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"hello": "world"})
	})
}

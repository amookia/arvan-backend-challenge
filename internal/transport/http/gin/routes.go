package gin

import "github.com/gin-gonic/gin"

func (a api) routes(c *gin.RouterGroup) {
	c.POST("/hello", a.middleware.requestQuota, func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"hello": "world"})
	})

	object := c.Group("/object", a.middleware.requestQuota)
	{
		object.POST("/put", a.middleware.monthlyQuota, a.upload.Put)
	}
}

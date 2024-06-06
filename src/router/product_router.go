package router

import (
	"github.com/gin-gonic/gin"
)

func ProductRouter(route *gin.RouterGroup) {
	route.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"message": "GET product by ID",
			"id":      id,
		})
	})
}

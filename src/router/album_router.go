package router

import (
	"GO/src/controllers"

	"github.com/gin-gonic/gin"
)

func AlbumRouter(route *gin.RouterGroup) {
	route.GET("/", controllers.GetAlbums)
	route.GET("/albums/:id", controllers.GetAlbumByID)

	route.POST("/albums", controllers.PostAlbums)

	route.DELETE("/albums/:id", controllers.DeleteAlbumByID)

	route.PUT("/albums/:id", controllers.UpdateAlbumByID)
}

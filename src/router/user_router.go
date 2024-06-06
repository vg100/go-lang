package router

import (
	"GO/src/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(route *gin.RouterGroup) {

	route.POST("/login", controllers.Login)
	route.POST("/register", controllers.Register)

}

package Router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	//v1 := router.Group("information").Use(Middlewares.AuthMiddleware).Use(Middlewares.Cors)
	//{
	//	v1.POST("/create", Controllers.WzhCreate)
	//}

	router.Run(":8080")
}

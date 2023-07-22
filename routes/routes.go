package routes

import (
	"github.com/gin-gonic/gin"
	"iae.com/smartpower/handlers"
)

func LoadRoutes(r *gin.Engine) {
	// 定义GET请求
	r.GET("/", handlers.HomeHandler)
	r.GET("/users/:id", handlers.GetUserByID)
}

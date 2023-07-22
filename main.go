package main

import (
	"github.com/gin-gonic/gin"
	"iae.com/smartpower/routes"
	"iae.com/smartpower/models"
)

func main() {
	// 初始化数据连接
	models.InitDB()

	// 创建默认引擎
	r := gin.Default()

	//加载路由
	routes.LoadRoutes(r)

	//启动
	r.Run(":9090")
}

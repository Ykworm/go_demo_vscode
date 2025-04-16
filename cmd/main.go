package main

import (
	"gin-demo-project/internal/handlers"
	"gin-demo-project/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化Gin引擎
	r := gin.Default()

	// 创建handler实例
	h := handlers.NewHandler()

	// 设置路由
	routes.SetupRoutes(r, h)

	// 启动服务
	r.Run() // 默认监听 0.0.0.0:8080
}

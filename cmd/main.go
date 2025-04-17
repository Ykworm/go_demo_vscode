package main

import (
	"database/sql"
	"gin-demo-project/internal/handlers"
	"gin-demo-project/internal/routes"

	"github.com/gin-gonic/gin"
)

func DatabaseMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

func main() {
	// 初始化Gin引擎
	r := gin.Default()
	db, _ := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/mydb")

	r.Use(DatabaseMiddleware(db))
	// 创建handler实例
	h := handlers.NewHandler()

	// 设置路由
	routes.SetupRoutes(r, h)

	// 启动服务
	r.Run() // 默认监听 0.0.0.0:8080
}

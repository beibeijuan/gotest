// main.go
package main

import (
	"goproject/database"
	"goproject/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect() // 连接数据库

	router := gin.Default()

	// 设置路由
	router.POST("/users", handlers.CreateUser)
	router.GET("/users", handlers.GetAllUsers)
	router.PUT("/users/:id", handlers.UpdateUser)
	router.DELETE("/users/:id", handlers.DeleteUser)

	// 启动服务器
	router.Run(":8080")
}

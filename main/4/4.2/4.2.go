package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	//// 创建一个Gin实例
	//router := gin.Default()
	//
	//// 定义一个路由和处理函数
	//router.GET("/hello", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "Hello, Gin!",
	//	})
	//})
	//
	//// 启动Gin服务器
	//router.Run(":8080")

	router := gin.Default()

	// 使用中间件
	router.Use(Logger())

	// 定义路由
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin!",
		})
	})

	router.GET("/goodbye", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Goodbye, Gin!",
		})
	})

	router.Run(":8080")
}

// 定义一个中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 在处理请求之前执行的操作
		fmt.Println("Logger Middleware - Before Request")
		start := time.Now()

		// 处理请求
		c.Next()
		elapsed := time.Since(start)
		fmt.Printf("Request processed in %s\n", elapsed)

		// 在处理请求之后执行的操作
		fmt.Println("Logger Middleware - After Request")
	}
}

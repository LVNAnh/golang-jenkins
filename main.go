package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Khởi tạo một router mới
	router := gin.Default()

	// Định nghĩa một route GET cơ bản
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Gin Gonic running on port 3000!",
		})
	})

	// Định nghĩa một route khác cho kiểm tra
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Chạy server trên cổng 8080
	router.Run(":3000") // Mặc định chạy trên localhost:8080
}

package main

import "github.com/gin-gonic/gin"

func main() {
	// database.gormConnect()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "200"})
	})
	r.Run(":8080")
}

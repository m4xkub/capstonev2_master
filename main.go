package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/PrimaryIP", func(c *gin.Context) {

		c.JSON(200, gin.H{"primary_ip": ""})
	})

	r.Run(":8080")
}

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/m4xkub/capstonev2_master/services"
)

func main() {
	r := gin.Default()

	//Add node
	r.POST("/AddNode", services.AddNode)

	// Get primary IP
	r.GET("/PrimaryIP", services.GetPrimary)

	// Check if this node is primary
	r.GET("/HavePrimary", services.HavePrimary)

	// Update primary
	r.GET("/UpdatePrimary", services.UpdateCurrentPrimary)

	r.Run(":8080")
}

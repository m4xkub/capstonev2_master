package route

import (
	"github.com/gin-gonic/gin"
	"github.com/m4xkub/capstonev2_master/services"
)

func Route() {
	r := gin.Default()

	//Add node
	r.POST("/AddNode", services.AddNode)

	// Get primary IP
	r.GET("/PrimaryIP", services.GetPrimary)

	// Check if this node is primary
	r.GET("/HavePrimary", services.HavePrimary)

	// Update primary
	r.GET("/UpdatePrimary", services.UpdateCurrentPrimary)

	r.GET("/EnableCluster1", services.EnableCluster1)
	r.GET("/DestroyCluster", services.DestroyCluster)
	r.GET("/EnableCluster2", services.EnableCluster2)
	r.GET("/MigrateCluster", services.MigrateCluster)

	r.Run(":8080")
}

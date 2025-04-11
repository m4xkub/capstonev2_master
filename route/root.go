package route

import (
	"github.com/gin-gonic/gin"
	"github.com/m4xkub/capstonev2_master/services"
	terraformservice "github.com/m4xkub/capstonev2_master/services/TerraformService"
)

func Route() {
	r := gin.Default()

	// Get primary IP
	r.GET("/PrimaryIP", services.GetPrimary)

	// Check if this node is primary
	r.GET("/HavePrimary", services.HavePrimary)

	// Update primary
	r.GET("/UpdatePrimary", services.UpdateCurrentPrimary)

	r.GET("/EnableCluster1", terraformservice.EnableCluster1)
	r.GET("/EnableCluster2", terraformservice.EnableCluster2)
	r.GET("/DestroyCluster", terraformservice.DestroyCluster)
	r.GET("/Enable2Cluster", terraformservice.Enable2Cluster)

	r.Run(":8080")
}

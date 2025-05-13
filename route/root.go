package route

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/m4xkub/capstonev2_master/services"
	apiservice "github.com/m4xkub/capstonev2_master/services/ApiService"
	handleterraform "github.com/m4xkub/capstonev2_master/services/HandleTerraform"
	terraformservice "github.com/m4xkub/capstonev2_master/services/TerraformService"
)

func Route() {
	r := gin.Default()

	// Get primary IP
	r.POST("/PrimaryIP", services.GetPrimary)

	// Check if this node is primary
	r.GET("/HavePrimary", services.HavePrimary)

	// Update primary
	r.GET("/UpdatePrimary", services.UpdateCurrentPrimary)

	r.GET("/EnableCluster1", terraformservice.EnableCluster1)
	r.GET("/EnableCluster2", terraformservice.EnableCluster2)
	r.GET("/DestroyCluster", terraformservice.DestroyCluster)
	r.GET("/Enable2Cluster", terraformservice.Enable2Cluster)

	r.GET("/test", func(c *gin.Context) {
		x, err := apiservice.Get("http://43.209.5.156:8080/healthCheck")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println((*x)["role"])
		fmt.Println((*x)["disk-status"])

		c.JSON(200, *x)
	})

	r.GET("/testMigrate", func(c *gin.Context) {
		handleterraform.MigrateData()
	})

	// r.GET("testdrbd", func(c *gin.Context) {
	// 	drbdservice.InitDrbd(cluster.DiskCluster)
	// })

	r.Run(":8080")
}

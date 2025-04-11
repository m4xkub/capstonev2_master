package terraformservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m4xkub/capstonev2_master/classes/terraform"
	handleterraform "github.com/m4xkub/capstonev2_master/services/HandleTerraform"
)

func EnableCluster1(c *gin.Context) {

	terraform.EnableCluster1()
	c.JSON(http.StatusOK, gin.H{
		"message": "Cluster 1 enabled",
	})

	handleterraform.HandleTerraformCluster()
}

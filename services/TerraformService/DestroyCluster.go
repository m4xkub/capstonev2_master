package terraformservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m4xkub/capstonev2_master/classes/terraform"
	handleterraform "github.com/m4xkub/capstonev2_master/services/HandleTerraform"
)

func DestroyCluster(c *gin.Context) {
	terraform.DestroyCluster()

	c.JSON(http.StatusOK, gin.H{
		"message": "Cluster destroy",
	})
	handleterraform.HandleTerraformCluster()
}

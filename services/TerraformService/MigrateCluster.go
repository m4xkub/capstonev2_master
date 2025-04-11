package terraformservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m4xkub/capstonev2_master/classes/terraform"
	handleterraform "github.com/m4xkub/capstonev2_master/services/HandleTerraform"
)

func Enable2Cluster(c *gin.Context) {

	terraform.Enable2Cluster()
	c.JSON(http.StatusOK, gin.H{
		"message": "migrate cluster",
	})

	// call function in HandleMigrate
	handleterraform.HandleTerraformCluster()
}

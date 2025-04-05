package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m4xkub/capstonev2_master/classes/terraform"
)

func MigrateCluster(c *gin.Context) {

	terraform.MigrateCluster()
	c.JSON(http.StatusOK, gin.H{
		"message": "migrate cluster",
	})
}

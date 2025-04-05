package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m4xkub/capstonev2_master/classes/terraform"
)

func DestroyCluster(c *gin.Context) {
	terraform.DestroyCluster()

	c.JSON(http.StatusOK, gin.H{
		"message": "Cluster destroy",
	})
}

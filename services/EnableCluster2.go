package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m4xkub/capstonev2_master/classes/terraform"
)

func EnableCluster2(c *gin.Context) {

	terraform.EnableCluster2()
	c.JSON(http.StatusOK, gin.H{
		"message": "Cluster 2 enabled",
	})
}

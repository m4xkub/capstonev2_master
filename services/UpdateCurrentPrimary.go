package services

import (
	"github.com/gin-gonic/gin"
	"github.com/m4xkub/capstonev2_master/classes/cluster"
)

func UpdateCurrentPrimary(c *gin.Context) {
	err := cluster.ClusterInstance.UpdateCurrentPrimary()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Primary updated"})
}
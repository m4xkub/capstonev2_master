package services

import (
	"github.com/gin-gonic/gin"
	"github.com/m4xkub/capstonev2_master/classes/cluster"
)

func HavePrimary(c *gin.Context) {
	hp := cluster.ClusterInstance.HavePrimary()
	c.JSON(200, gin.H{"have_primary": hp})
}
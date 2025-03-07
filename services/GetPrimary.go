package services

import (
	"github.com/gin-gonic/gin"
	"github.com/m4xkub/capstonev2_master/classes/cluster"
)


func GetPrimary(c *gin.Context) {
	ip, err := cluster.ClusterInstance.GetPrimary()
	if err != nil { 
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"primary_ip": ip})
}
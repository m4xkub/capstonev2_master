package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m4xkub/capstonev2_master/classes/cluster"
	"github.com/m4xkub/capstonev2_master/classes/node"
)

type AddNodeRequest struct {
	NodeIpAddress string `json:"node_ip_address" binding:"required"`
}

func AddNode(c *gin.Context) {
	var requestBody AddNodeRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cluster.ClusterInstance.NodesInCluster = append(cluster.ClusterInstance.NodesInCluster, node.NewNode(requestBody.NodeIpAddress))

	c.JSON(http.StatusOK, gin.H{
		"message": "Node added successfully",
		"data":    requestBody,
	})
}

package services

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/m4xkub/capstonev2_master/classes/cluster"
	apiservice "github.com/m4xkub/capstonev2_master/services/ApiService"
)

type GetPrimaryRequest struct {
	PublicIp string `json:"ip"`
}

func GetPrimary(c *gin.Context) {

	var req GetPrimaryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	data := make(map[string]interface{})

	data["ip"] = req.PublicIp
	ip, err := cluster.ClusterInstance.GetPrimary()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	apiservice.Post(fmt.Sprintf("http://%s:8080/addClient", ip), &data)
	c.JSON(200, gin.H{"primary_ip": ip})
}

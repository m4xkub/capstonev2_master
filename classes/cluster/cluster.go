package cluster

import (
	"errors"
	"fmt"

	"github.com/m4xkub/capstonev2_master/classes/node"
	apiservice "github.com/m4xkub/capstonev2_master/services/ApiService"
)

type Cluster struct {
	IsInitDrbd     bool
	CurrentPrimary *node.Node
	NodesInCluster []*node.Node
}

var ClusterInstance *Cluster

var DiskCluster *Cluster
var MigratedDiskCluster *Cluster

func (c *Cluster) UpdateCurrentPrimary() error {
	if len(c.NodesInCluster) == 0 {
		return errors.New("no node in cluster")
	}

	for _, node := range c.NodesInCluster {
		status, err := node.CheckStatus()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if status.Role == "Primary" {
			c.CurrentPrimary = node
			return nil
		}
	}

	for _, node := range c.NodesInCluster {
		status, err := node.CheckStatus()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if status.Role == "Secondary" {
			node.PromoteToPrimary()
			c.CurrentPrimary = node

			return nil
		}
	}

	return nil

}

func (c *Cluster) HavePrimary() bool {

	for _, node := range c.NodesInCluster {
		status, err := node.CheckStatus()

		if err != nil {
			return false
		}

		if status.Role == "Primary" {
			return true
		}
	}

	print(len(c.NodesInCluster))

	return false
}

func (c *Cluster) GetPrimary() (string, error) {
	// if c.CurrentPrimary == nil {
	// 	return "", errors.New("update primary in cluster is needed before get one")
	// }
	// status, err := c.CurrentPrimary.CheckStatus()

	// if err != nil {
	// 	return "", err
	// }

	// if !(status.Role == "Primary") {
	// 	c.UpdateCurrentPrimary()
	// }
	// if !(status.DiskStatus == "UpToDate") {
	// 	return "", errors.New("waiting for primary to be ready")
	// }

	c.UpdateCurrentPrimary()
	apiservice.Get(fmt.Sprintf("http://%s:8080/initNbd", c.CurrentPrimary.PublicIp))

	return c.CurrentPrimary.PublicIp, nil
}

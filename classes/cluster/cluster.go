package cluster

import (
	"errors"
	"fmt"

	"github.com/m4xkub/capstonev2_master/classes/node"
)

type Cluster struct {
	CurrentPrimary *node.Node
	NodesInCluster []*node.Node
	//Status         string
}

var ClusterInstance Cluster

func (c *Cluster) UpdateCurrentPrimary() error {
	if len(c.NodesInCluster) == 0 {
		return errors.New("no node in cluster")
	}

	for _, node := range c.NodesInCluster {
		status, err := node.CheckStatus()
		if err != nil {
			return err
		}
		//&& status.DiskStatus == "UpToDate"
		if status.Role == "Primary" {
			c.CurrentPrimary = node
			return nil
		}
	}
	target_node := c.NodesInCluster[0]
	target_node.PromoteToPrimary()
	c.CurrentPrimary = target_node

	return nil

}

func (c *Cluster) HavePrimary() bool {

	for _, node := range c.NodesInCluster {
		fmt.Println("here")
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
	if c.CurrentPrimary == nil {
		return "", errors.New("update primary in cluster is needed before get one")
	}
	status, err := c.CurrentPrimary.CheckStatus()

	if err != nil {
		return "", err
	}

	if !(status.Role == "Primary") {
		c.UpdateCurrentPrimary()
	}
	if !(status.DiskStatus == "UpToDate") {
		return "", errors.New("waiting for primary to be ready")
	}
	return c.CurrentPrimary.IpAddress, nil
}

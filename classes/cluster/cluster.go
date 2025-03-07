package cluster

import (
	"errors"

	"github.com/m4xkub/capstonev2_master/classes/node"
)

type Cluster struct {
	CurrentPrimary *node.Node
	NodesInCluster []*node.Node
	//Status         string
}

func (c *Cluster) FindPrimary() error {
	if len(c.NodesInCluster) == 0 {
		return errors.New("no node in cluster")
	}

	for _, node := range c.NodesInCluster {
		status, err := node.CheckStatus()
		if err != nil {
			return err
		}

		if status.Role == "Primary" && status.DiskStatus == "UpToDate" {
			c.CurrentPrimary = node
			return nil
		}
	}

	return errors.New("primary not found")
}

func (c *Cluster) PromoteNewPrimary(ipaddr string) error {

	return nil
}

func (c *Cluster) HavePrimary() bool {

	for _, node := range c.NodesInCluster {
		status, _ := node.CheckStatus()

		if status.Role == "Primary" {
			return true
		}
	}

	return false
}

func (c *Cluster) GetPrimary() (string, error) {
	status, err := c.CurrentPrimary.CheckStatus()
	if err != nil {
		return "", err
	}

	if !(status.Role == "Primary" && status.DiskStatus == "UpToDate") {
		c.FindPrimary()
	}
	return c.CurrentPrimary.IpAddress, nil
}

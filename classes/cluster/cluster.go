package cluster

import (
	"github.com/m4xkub/capstonev2_master/classes/node"
)

type Cluster struct {
	CurrentPrimary *node.Node
	NodesInCluster []*node.Node
	Status         string
}

func (c *Cluster) UpdateNodesInCluster() {
	var NewCluster []*node.Node
	primaryIsUp := true
	for _, n := range c.NodesInCluster {
		okay := n.CheckStatus()

		if okay {
			NewCluster = append(NewCluster, n)
			continue
		}

		if n.IpAddress == c.CurrentPrimary.IpAddress {
			primaryIsUp = false
		}
	}

	if !primaryIsUp {
		if len(NewCluster) == 0 {
			// add more node before promote
			for range 3 {
				NewCluster = append(NewCluster, node.NewNode())
			}
		}

		// promote new primary
		c.PromoteNewPrimary(NewCluster[0].IpAddress)
		c.CurrentPrimary = NewCluster[0]
	}

	c.NodesInCluster = NewCluster
}

func (c *Cluster) PromoteNewPrimary(ipaddr string) error {
	//to do
	return nil
}

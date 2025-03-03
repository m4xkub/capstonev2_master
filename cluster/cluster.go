package cluster

import "net/http"

type Cluster struct {
	CurrentPrimary string
	NodesInCluster []string
	Status         string
}

func (c *Cluster) CheckStatus(ip string) bool {
	_, err := http.Get(c.CurrentPrimary + "/healthCheck")

	if err != nil {
		return false
	} else {
		return true
	}
}

func (c *Cluster) UpdateNodesInCluster() {
	var NewCluster []string
	primaryIsUp := true
	for _, ipaddr := range c.NodesInCluster {
		okay := c.CheckStatus(ipaddr)

		if okay {
			NewCluster = append(NewCluster, ipaddr)
			continue
		}

		if ipaddr == c.CurrentPrimary {
			primaryIsUp = false
		}
	}

	if !primaryIsUp {
		if len(NewCluster) == 0 {
			// add more node before promote
		}

		// promote new primary
		c.PromoteNewPrimary(NewCluster[0])
		c.CurrentPrimary = NewCluster[0]
	}

	c.NodesInCluster = NewCluster
}

func (c *Cluster) PromoteNewPrimary(ipaddr string) error {
	//to do
	return nil
}

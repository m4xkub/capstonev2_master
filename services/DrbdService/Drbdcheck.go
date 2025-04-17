package drbdservice

import (
	"fmt"

	"github.com/m4xkub/capstonev2_master/classes/cluster"
	apiservice "github.com/m4xkub/capstonev2_master/services/ApiService"
)

func IsInitedDrbd(cluster *cluster.Cluster) bool {
	for _, e := range cluster.NodesInCluster {
		_, err := apiservice.Get(fmt.Sprintf("http://%s:8080/drbdCheck", e.PublicIp))
		if err != nil {
			return false
		}
	}

	return true
}

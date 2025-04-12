package drbdservice

import (
	"fmt"

	"github.com/m4xkub/capstonev2_master/classes/cluster"
	apiservice "github.com/m4xkub/capstonev2_master/services/ApiService"
)

func InitDrbd() {
	data := make(map[string]interface{})

	data["private_ip_1"] = cluster.ClusterInstance.NodesInCluster[0].PrivateIp
	data["private_ip_2"] = cluster.ClusterInstance.NodesInCluster[1].PrivateIp
	data["disk_name"] = "/dev/nvme1n1"

	for _, e := range cluster.ClusterInstance.NodesInCluster {
		apiservice.Post(fmt.Sprintf("http://%s:8080/initConfigFile", e.PublicIp), &data)
		apiservice.Post(fmt.Sprintf("http://%s:8080/initMetaData", e.PublicIp), &data)
	}
}

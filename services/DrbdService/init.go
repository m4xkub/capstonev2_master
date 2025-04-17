package drbdservice

import (
	"fmt"
	"net/http"
	"time"

	"github.com/m4xkub/capstonev2_master/classes/cluster"
	apiservice "github.com/m4xkub/capstonev2_master/services/ApiService"
)

func InitDrbd(cluster *cluster.Cluster) {

	data := make(map[string]interface{})

	data["private_ip_1"] = cluster.NodesInCluster[0].PrivateIp
	data["private_ip_2"] = cluster.NodesInCluster[1].PrivateIp
	data["disk_name"] = "/dev/nvme1n1"

	for {
		ready := true
		for _, e := range cluster.NodesInCluster {

			resp, err := http.Get(fmt.Sprintf("http://%s:8080/healthCheck", e.PublicIp))

			if err != nil {
				ready = false
				fmt.Println(err.Error())
				break
			}

			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				ready = false
				fmt.Println("unexpected status code: %d", resp.StatusCode)
				break
			}
		}

		if ready {
			break
		}
		time.Sleep(5 * time.Second)
	}

	for _, e := range cluster.NodesInCluster {
		apiservice.Post(fmt.Sprintf("http://%s:8080/initConfigFile", e.PublicIp), &data)
		apiservice.Post(fmt.Sprintf("http://%s:8080/initMetaData", e.PublicIp), &data)
	}

	// promote to primary
	apiservice.Get(fmt.Sprintf("http://%s:8080/promote", cluster.NodesInCluster[0].PublicIp))

	for {
		// wait for two disk to sync
		res, err := apiservice.Get(fmt.Sprintf("http://%s:8080/healthCheck", cluster.NodesInCluster[0].PublicIp))
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if (*res)["disk-status"] == "UpToDate" {
			break
		}
		time.Sleep(5 * time.Second)
	}

	// make filesystem
	apiservice.Get(fmt.Sprintf("http://%s:8080/makeFileSystem", cluster.NodesInCluster[0].PublicIp))

	//mount file
	apiservice.Get(fmt.Sprintf("http://%s:8080/mountVolume", cluster.NodesInCluster[0].PublicIp))
}

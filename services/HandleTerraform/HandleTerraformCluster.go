package handleterraform

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/m4xkub/capstonev2_master/classes/cluster"
	"github.com/m4xkub/capstonev2_master/classes/node"
)

type OutputEntry struct {
	Sensitive bool          `json:"sensitive"`
	Type      []interface{} `json:"type"`
	Value     []string      `json:"value"`
}

type OutputData map[string]OutputEntry

func HandleTerraformCluster() {
	//read output.json
	file, err := os.Open("./output.json")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var data OutputData

	byteValue, _ := io.ReadAll(file)
	json.Unmarshal([]byte(byteValue), &data)

	if len(data) == 0 {
		fmt.Println("JSON file is empty or just contains {}")
		cluster.ClusterInstance = nil
		cluster.DiskCluster = nil
		cluster.MigratedDiskCluster = nil
		return
	}

	// write data to cluster package
	haveDiskCluster := false
	haveMigrateDiskCluster := false

	if len(data["disk_private_ips"].Value) != 0 {
		var tmpNodes []*node.Node
		var tmpDiskCluster cluster.Cluster

		for i := range 3 {
			var n node.Node
			n.PrivateIp = data["disk_private_ips"].Value[i]
			n.PublicIp = data["disk_public_ips"].Value[i]
			tmpNodes = append(tmpNodes, &n)
		}

		tmpDiskCluster.NodesInCluster = tmpNodes
		cluster.DiskCluster = &tmpDiskCluster
		haveDiskCluster = true
	}

	if len(data["disk_migrate_private_ips"].Value) != 0 {
		var tmpNodes []*node.Node
		var tmpMigrateDiskCluster cluster.Cluster

		for i := range 3 {
			var n node.Node
			n.PrivateIp = data["disk_migrate_private_ips"].Value[i]
			n.PublicIp = data["disk_migrate_public_ips"].Value[i]
			tmpNodes = append(tmpNodes, &n)
		}

		tmpMigrateDiskCluster.NodesInCluster = tmpNodes
		cluster.MigratedDiskCluster = &tmpMigrateDiskCluster

		haveMigrateDiskCluster = true
	}

	if haveDiskCluster && haveMigrateDiskCluster {
		// do something
		MigrateData()
	} else if haveDiskCluster {
		cluster.ClusterInstance = cluster.DiskCluster
	} else if haveMigrateDiskCluster {
		cluster.ClusterInstance = cluster.MigratedDiskCluster
	} else {
		// didnt have any cluster
		cluster.ClusterInstance = nil
	}
	fmt.Println("**********************")
	for idx, e := range cluster.ClusterInstance.NodesInCluster {
		fmt.Println(idx, *e)
	}
	fmt.Println("**********************")

}

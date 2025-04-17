package handleterraform

import (
	"fmt"

	"github.com/m4xkub/capstonev2_master/classes/cluster"
)

func MigrateData() {
	start := cluster.ClusterInstance
	stop := cluster.DiskCluster
	if start == stop {
		stop = cluster.MigratedDiskCluster
	}

	_, err := stop.GetPrimary()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

}

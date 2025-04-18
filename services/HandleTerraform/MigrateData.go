package handleterraform

import (
	"fmt"
	"reflect"

	"github.com/m4xkub/capstonev2_master/classes/cluster"
	apiservice "github.com/m4xkub/capstonev2_master/services/ApiService"
)

func MigrateData() {
	start := cluster.ClusterInstance
	stop := cluster.DiskCluster
	if reflect.DeepEqual(start, stop) {
		stop = cluster.MigratedDiskCluster
	}

	start_ip, err := start.GetPrimary()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	stop_ip, err := stop.GetPrimary()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	data := make(map[string]interface{})
	// *** ip got from GetPrimary function is public_ip this would still work find but can be fix for more level in security
	data["private_ip"] = stop_ip
	apiservice.Post(fmt.Sprintf("http://%s:8080/migrate", start_ip), &data)

}

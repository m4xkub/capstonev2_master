package handleterraform

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/m4xkub/capstonev2_master/classes/cluster"
)

func WaitForInstance() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-southeast-7"))
	if err != nil {
		fmt.Println("err here")
		panic(err)
	}
	client := ec2.NewFromConfig(cfg)

	if cluster.ClusterInstance == nil {
		return
	}

	for _, e := range cluster.ClusterInstance.NodesInCluster {
		waitUntilRunning(client, e.Id)
		waitUntilStatusOK(client, e.Id)
	}

	fmt.Println("All instances are ready!")

}

func waitUntilRunning(client *ec2.Client, instanceID string) {
	for {
		out, err := client.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{
			InstanceIds: []string{instanceID},
		})
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		state := out.Reservations[0].Instances[0].State.Name
		fmt.Println("Instance state:", state)
		if state == "running" {
			break
		}
		time.Sleep(5 * time.Second)
	}
}

func waitUntilStatusOK(client *ec2.Client, instanceID string) {
	for {
		status, err := client.DescribeInstanceStatus(context.TODO(), &ec2.DescribeInstanceStatusInput{
			InstanceIds: []string{instanceID},
		})
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		if len(status.InstanceStatuses) > 0 {
			inst := status.InstanceStatuses[0]
			if inst.InstanceStatus.Status == "ok" && inst.SystemStatus.Status == "ok" {
				break
			}
		}

		fmt.Println("Waiting for instance status checks to pass...")
		time.Sleep(5 * time.Second)
	}
}

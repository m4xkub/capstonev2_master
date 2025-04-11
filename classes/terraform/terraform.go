package terraform

import (
	"fmt"
	"os"
	"os/exec"
)

func InitTerraform() {
	_ = exec.Command("terraform", "-chdir=terraform", "init")
}

func Enable2Cluster() {
	createClusterVariableFile := "-var-file=migrate.tfvars"
	cmd := exec.Command("terraform", "-chdir=terraform", "apply", createClusterVariableFile, "-auto-approve")

	// Run and get output
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
	}

	writeOutputFile()
}

func EnableCluster1() {
	createClusterVariableFile := "-var-file=enable_cluster_1.tfvars"
	cmd := exec.Command("terraform", "-chdir=terraform", "apply", createClusterVariableFile, "-auto-approve")

	// Run and get output
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
	}

	writeOutputFile()

}

func EnableCluster2() {
	createClusterVariableFile := "-var-file=enable_cluster_2.tfvars"
	cmd := exec.Command("terraform", "-chdir=terraform", "apply", createClusterVariableFile, "-auto-approve")

	// Run and get output
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
	}

	writeOutputFile()

}

func DestroyCluster() {
	createClusterVariableFile := "-var-file=migrate.tfvars"
	cmd := exec.Command("terraform", "-chdir=terraform", "destroy", createClusterVariableFile, "-auto-approve")

	// Run and get output
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
	}

	writeOutputFile()
}

func writeOutputFile() {
	cmd := exec.Command("terraform", "-chdir=terraform", "output", "-json")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = os.WriteFile("output.json", output, 0644)
	if err != nil {
		fmt.Println("Write error:", err)
		return
	}

	fmt.Println("✅ output.json written successfully!")
}

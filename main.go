package main

import (
	"fmt"
	"os/exec"

	"github.com/m4xkub/capstonev2_master/route"
	handleterraform "github.com/m4xkub/capstonev2_master/services/HandleTerraform"
)

func main() {
	cmd := exec.Command("terraform", "-chdir=terraform", "init")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Output:", string(output))
	} else {
		fmt.Println("Success:", string(output))
	}
	handleterraform.HandleTerraformCluster()
	route.Route()
}

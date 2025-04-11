package main

import (
	"fmt"
	"os/exec"

	"github.com/m4xkub/capstonev2_master/route"
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
	route.Route()
}

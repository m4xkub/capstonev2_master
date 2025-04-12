package handleterraform

import (
	"fmt"
	"os/exec"
)

func MigrateData() {
	cmd := exec.Command("terraform")

	// Run and get output
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

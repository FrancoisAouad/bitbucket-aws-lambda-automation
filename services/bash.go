package services

import (
	"example/lib"
	"os"
	"os/exec"
)

var logger=lib.Log

func ExecShell() {
	// call event bus type handler here
	// pass array to string
	cmd := exec.Command("bash", "scripts/update_branches.sh")
	_, err := cmd.Output()
	if err != nil {
		logger.WithError(err).Error("Failed to execute bash script")
		os.Exit(1)
	}
}

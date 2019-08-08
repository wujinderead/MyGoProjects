package util

import (
	"fmt"
	"os/exec"
)

func RunCmd(command string, args ...string) {
	cmd := exec.Command(command, args...)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("exec err:", err.Error())
	}
	fmt.Println(string(output))
}

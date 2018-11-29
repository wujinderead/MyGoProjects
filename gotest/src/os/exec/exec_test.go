package exec

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestLookPath(t *testing.T) {
	for _, command := range []string{"ls", "python", "python3", "java"} {
		path, err := exec.LookPath(command)
		if err != nil {
			fmt.Printf("cmd: %s, err: %s\n", command, err.Error())
		}
		fmt.Printf("%s: %s\n", command, path)
	}
}

func Test(t *testing.T) {

}

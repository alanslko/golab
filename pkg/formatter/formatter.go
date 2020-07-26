package formatter

import (
	"bytes"
	"fmt"
	"os/exec"
)

func GetName() string {
	cmd := exec.Command("script/test.sh")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		fmt.Printf("cmd.Run() failed with %s\n", err)
	}

	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)

	return "formatter"
}

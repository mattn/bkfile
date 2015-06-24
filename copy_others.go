// +build !windows

package main

import (
	"fmt"
	"os/exec"
)

func copyFile(src, dst string) error {
	b, err := exec.Command("/bin/cp", "-n", src, dst).CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s: %s", err, string(b))
	}
	return nil
}

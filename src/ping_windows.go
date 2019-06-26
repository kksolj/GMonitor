// +build windows
package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func doPing(addr string, count uint, timeout uint) bool {
	_, err := exec.LookPath("ping")
	if err != nil {
		return GoPing(addr, count, timeout)
	}
	out, _ := exec.Command("ping", addr, "-n", fmt.Sprintf("%i", count), "-w", fmt.Sprintf("%i", timeout)).Output()
	return strings.Contains(string(out), "Destination Host Unreachable")
}

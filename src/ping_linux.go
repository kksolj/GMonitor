// +build linux
package main

import (
	"fmt"
	"os/exec"
)

func doPing(addr string, count uint, timeout uint) bool {
	_, err := exec.LookPath("ping")
	if err != nil {
		return GoPing(addr, count, timeout)
	}
	_, er := exec.Command("ping", addr, "-c", fmt.Sprintf("%v", count), "-W", fmt.Sprintf("%v", timeout)).Output()
	return er == nil
}

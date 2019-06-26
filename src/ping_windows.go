// +build windows
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
	_, er := exec.Command("ping", fmt.Sprintf(` %s -n %v -w %v`, addr, count, timeout)).Output()
	//log.Printf("ping status %v =>%s", addr, er, er != nil)
	return er == nil
}

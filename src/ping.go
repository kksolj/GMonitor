package main

import (
	"github.com/sparrc/go-ping"
	"time"
)

func GoPing(addr string, count uint, timeout uint) bool {
	pinger, err := ping.NewPinger(addr)
	if err != nil {
		return false
	}
	pinger.SetPrivileged(true)
	pinger.Interval = time.Millisecond * time.Duration(timeout+1)
	pinger.Timeout = time.Millisecond * time.Duration(timeout)
	pinger.Count = int(count)
	pinger.Run()                 // blocks until finished
	stats := pinger.Statistics() // get send/receive/rtt stats
	return stats.PacketLoss != 100
}

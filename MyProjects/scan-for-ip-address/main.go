package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-ping/ping"
)

func pingAddress(ip string, wg *sync.WaitGroup) {
	defer wg.Done()
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		fmt.Println("Ping error:", err)
		return
	}
	pinger.Count = 3
	pinger.Timeout = time.Second * 3
	pinger.SetPrivileged(true) // Note: might require running with sudo

	err = pinger.Run() // Blocks until finished.
	if err != nil {
		fmt.Println("Ping execution error for ip:", ip, err)
		return
	}

	stats := pinger.Statistics() // get send/receive/duplicate/rtt stats
	if stats.PacketsRecv > 0 {
		fmt.Println("Active IP:", ip)
	}
}

func main() {
	var wg sync.WaitGroup

	// Replace "192.168.1." with the prefix of your local network IP range
	for i := 1; i <= 250; i++ {
		ip := fmt.Sprintf("10.2.9.%d", i)
		wg.Add(1)
		go pingAddress(ip, &wg)
	}

	wg.Wait()
	fmt.Println("Finished scanning.")
}

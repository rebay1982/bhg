package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)

		if err != nil {
			results <- 0
			continue
		}

		conn.Close()
		results <- p
	}
}

func main() {

	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	// Create the 100 worker pool.
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	// Send all the work into the channel.
	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-results

		if port != 0 {
			openports = append(openports, port)
		}

		progress := (i * 100) / 1023
		fmt.Printf("Progress %d (%d)\n", progress, i)
	}

	close(ports)
	close(results)

	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("Port [%d] is open.\n", port)
	}
}


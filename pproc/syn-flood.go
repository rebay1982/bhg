package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	snaplen  = int32(320)
	promisc  = true
	timeout  = pcap.BlockForever
	filter   = "(tcp[13] == 0x11 or tcp[13] == 0x10 or tcp[13] == 0x18) and src host %s and src port %s"
	devFound = false
	results  = make(map[string]int)

	lock   = sync.RWMutex{}
	notify = make(chan empty)
)

func capture(iface, target string, port string) {
	handle, err := pcap.OpenLive(iface, snaplen, promisc, timeout)
	if err != nil {
		log.Panicln(err)
	}
	defer handle.Close()

	parsedFilter := fmt.Sprintf(filter, target, port)
	if err := handle.SetBPFFilter(parsedFilter); err != nil {
		log.Panicln(err)
	}

	source := gopacket.NewPacketSource(handle, handle.LinkType())
	for range source.Packets() {

		lock.Lock()
		results[port] += 1
		lock.Unlock()
	}
}

type empty struct{}

func main() {

	if len(os.Args) != 4 {
		log.Fatalln("Usage: main.go <capture_iface> <target_ip> <port1,port2,port3>")
	}

	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Panicln(err)
	}

	iface := os.Args[1]
	for _, device := range devices {
		if device.Name == iface {
			devFound = true
		}
	}
	if !devFound {
		log.Panicf("Device named '%s' does not exist\n", iface)
	}

	ip := os.Args[2]

	ports, err := explode(os.Args[3])
	if err != nil {
		log.Panicln(err)
	}

	fmt.Printf("Setting up port capture for ports [%s]\n\n", os.Args[3])
	for _, port := range ports {
		go capture(iface, ip, port)
	}
	time.Sleep(1 * time.Second)

	for _, port := range ports {
		go scanPort(notify, ip, port)
	}
	time.Sleep(2 * time.Second)

	// Gather
	for i := 0; i < len(ports); i++ {
		<-notify
	}

	fmt.Println("\n---[RESULTS]------------------------------------------")
	for port, confidence := range results {
		if confidence >= 1 {
			fmt.Printf("Port [%s] open (confidence: %d)\n", port, confidence)
		}
	}
}

func scanPort(notify chan empty, ip string, port string) {

	target := fmt.Sprintf("%s:%s", ip, port)

	fmt.Printf("Trying [%s]...\n", target)

	c, err := net.DialTimeout("tcp", target, 1000*time.Millisecond)
	if err == nil {
		c.Close()
	}

	// Notify
	var e empty
	notify <- e
}

func explode(portString string) ([]string, error) {
	ret := make([]string, 0)

	ports := strings.Split(portString, ",")
	for _, port := range ports {
		port := strings.TrimSpace(port)
		ret = append(ret, port)
	}

	return ret, nil
}

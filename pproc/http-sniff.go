package main

import (
	"fmt"
	"log"

	"encoding/hex"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	iface    = "wlp1s0"
	snaplen  = int32(1600)
	promisc  = false
	timeout  = pcap.BlockForever
	filter   = "tcp and port 80"
	devFound = false
)

func main() {
	devices, err := pcap.FindAllDevs()

	if err != nil {
		log.Panicln(err)
	}

	for _, device := range devices {
		if device.Name == iface {
			devFound = true
		}
	}

	if !devFound {
		log.Panicf("Device named [%s] cannot be found\n", iface)
	}

	handle, err := pcap.OpenLive(iface, snaplen, promisc, timeout)

	if err != nil {
		log.Panicln(err)
	}
	defer handle.Close()

	// Can't set the filter.
	if err := handle.SetBPFFilter(filter); err != nil {
		log.Panicln(err)
	}

	source := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range source.Packets() {
		fmt.Println(packet)

		for _, layer := range packet.Layers() {

			fmt.Println(layer.LayerType())
			fmt.Printf("\tContent:\n%s\n", hex.Dump(layer.LayerContents()))
			fmt.Printf("\tPayload:\n%s\n", hex.Dump(layer.LayerPayload()))

		}
	}
}

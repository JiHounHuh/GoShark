package Pcap

import (
	"fmt"
//	"os"
//	"time"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
//	"github.com/google/gopacket/layers"
)

func Demo (Dev string) {
	fmt.Printf("%s\n",Dev)
	// Open interface
	if handle, err := pcap.OpenLive(Dev, 1600, true, 0); err != nil {
		panic(err)
	} else if err := handle.SetBPFFilter("tc amd prt 80"); err != nil {
		panic(err)
	} else { // Run if no errors
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			// Process packets here
			fmt.Println(packet)
			//w.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
		}
	}
	// Start capturing network traffic
	//
}
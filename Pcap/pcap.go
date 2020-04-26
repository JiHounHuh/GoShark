package Pcap

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"os"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func Capture (Dev string) {
	fmt.Printf("%s\n",Dev)
	// Open interface
	if handle, err := pcap.OpenLive(Dev, 1600, true, 0); err != nil {
		panic(err)
	} else { // Run if no errors
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		count := 0
		for packet := range packetSource.Packets() {
			// Process packets here
			//fmt.Println(packet.Data())
			fmt.Println("Packet =", count)
			filename := ("Packet"+ strconv.Itoa(count)+".pcap")
			writeErr := ioutil.WriteFile(filename, packet.Data(), 0644)

			if writeErr != nil {
				fmt.Println("There was a problem writing the file, but we have printed to the screen anyway")
				os.Exit(1)
			}
			count += 1
		}
	}
	// Start capturing network traffic
	//
}

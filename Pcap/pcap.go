package Pcap

import (
	"fmt"
	"strings"
	"time"
	"strconv"
	"os"
	"os/exec"
	"io/ioutil"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func sendToGCP(filename string) {
	scp := "/usr/bin/scp"
	user := "citrus"

	// read server IP from file, so this way
	// our IP of our GCP computer engine isnt on github
	ip, readErr := ioutil.ReadFile("secretIP")

	if readErr != nil {
		fmt.Println("Error reading secret IP from file", readErr)
		os.Exit(1)
	}

	dst := string(ip)+":~/pcaps/" // replace with your IP of analyzing server
	// scp PacketX user@<IP>:~/pcaps/
	cmd := exec.Command(scp, filename, (user+"@"+dst))
	sendErr := cmd.Run()

	if sendErr != nil {
		fmt.Println("Cannot send to GCP", sendErr)
		os.Exit(1)
	}
}

func Capture (Dev string) {
	fmt.Printf("%s\n",Dev)
	// Open interface
	if handle, err := pcap.OpenLive(Dev, 1600, true, 0); err != nil {
		panic(err)
	} else { // Run if no errors
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		count := 0

		filename := "packets"+strconv.Itoa(count)+".pcap"
		file, err := os.OpenFile(filename, os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0660)

		if err != nil{
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(filename,"created.")

		defer file.Close()

		start := time.Now()
		fmt.Println("start time",start)
		for packet := range packetSource.Packets() {
			// Process packets here
			end := time.Now()
			elapsed := end.Sub(start)

			// time in nanoseconds
			if elapsed >= 30000000000 {
				count += 1
				file.Close()
				filename = "packets"+strconv.Itoa(count)+".pcap"
				file, err = os.OpenFile(filename, os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0660)

				if err != nil{
					fmt.Println(err)
					os.Exit(1)
				}
				fmt.Println("created new pcap file:",filename)
				start = time.Now()
				elapsed = 0
			}
			//fmt.Println(packet.Dump())

			layers := packet.Layers()

			if len(layers) < 4 {
				//fmt.Println("\nLESS THAN FOUR")
				continue
			}

			//fmt.Println("\n",gopacket.LayerString(layers[3]),"\n")
			//fmt.Println("\n",string(packet.ApplicationLayer().Payload()),"\n") //prints payload as a string

			layer2 := strings.Split(gopacket.LayerDump(layers[1]), " ")
			layer3 := strings.Split(gopacket.LayerDump(layers[2]), " ")

			if strings.Split(layer2[2],"=")[1] == "6" {
				//f*** you ipv6
				continue
			}

			if layer3[2] == "0" {
				continue
			}
			if layer3[3] == "0" {
				continue
			}

			srcIP := strings.Split(layer2[12],"=")
			dstIP := strings.Split(layer2[13],"=")
			srcPort := strings.Split(layer3[2],"=")
			dstPort := strings.Split(layer3[3],"=")

			fmt.Println("\nsrcIP = ",srcIP[1])
			fmt.Println("dstIP = ",dstIP[1])

			if srcPort[0] == "SrcPort" {
				fmt.Println("SrcPort = ",srcPort[1])
			}

			if dstPort[0] == "DstPort" {
				fmt.Println("DstPort = ",dstPort[1])
			}

			if dstIP[1] == "104.24.118.24" {
				fmt.Println("srcIP:",srcIP[1],"\ndstIP:",dstIP[1],"\nsrcPort:",srcPort[1],"\ndstPort:",dstPort[1])
				fmt.Println("\n",string(packet.ApplicationLayer().Payload()),"\n")
			}


			_, err := file.WriteString("\n\nstart\n\n")
			_, err = file.WriteString(packet.Dump())
			_, err = file.WriteString("\n\nend\n\n")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			//go sendToGCP(filename)
			//w.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
		}
	}
	// Start capturing network traffic
	//
}

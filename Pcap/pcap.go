package Pcap

import (
	A "../Latex"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"os"
	"strings"
)

func Capture(Dev string) {
	flag := 0
	count := 0
	keywords := []string{"admin", "Admin", "Set-Cookie", "cookie", "Cookie", "user", "User", "Pass", "pass", "password", "passwd", "Password", "Passwd", "key", "Key", "username", "Username"}
	// Open interface
	if handle, err := pcap.OpenLive(Dev, 1600, true, 0); err != nil {
		panic(err)
	} else { // Run if no errors
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

		filename := "toRead.txt"
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0660)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(filename, "created.")
		defer file.Close()
		defer A.MakeReport()
		for packet := range packetSource.Packets() {
			if flag == 0 {
				fmt.Println("Capture begins.")
				flag = 1
			}
			layers := packet.Layers()
			if len(layers) < 4 {
				continue
			}
			layer2 := strings.Split(gopacket.LayerDump(layers[1]), " ")
			layer3 := strings.Split(gopacket.LayerDump(layers[2]), " ")
			if strings.Split(layer2[2], "=")[1] == "6" {
				continue
			}
			if layer3[2] == "0" {
				continue
			}
			if layer3[3] == "0" {
				continue
			}
			srcIP := strings.Split(layer2[12], "=")
			dstIP := strings.Split(layer2[13], "=")
			srcPort := strings.Split(layer3[2], "=")
			dstPort := strings.Split(layer3[3], "=")
			if srcPort[1] == "80(http)" || dstPort[1] == "80(http)" {
				fmt.Println("Detected http traffic!")
				payload := string(packet.ApplicationLayer().Payload())
				outputToFile := ""
				payloadArr := strings.Split(payload, "\n")
				for _, v1 := range payloadArr {
					for _, v2 := range keywords {
						if strings.Contains(v1, v2) {
							outputToFile += v1[0:len(v1)-2] + " "
						}
					}
				}
				if outputToFile == "" {
					outputToFile = "Plaintext data recoverable.~\n"
				}
				line := srcIP[1] + "~" + dstIP[1] + "~" + srcPort[1] + "~" + dstPort[1] + "~" + outputToFile + "\n"
				_, err := file.WriteString(line)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}
			if (srcPort[1] == "20(ftp)" && dstPort[1] == "20(ftp)") || (srcPort[1] == "21(ftp)" && dstPort[1] == "21(ftp)") {
				line := srcIP[1] + "~" + dstIP[1] + "~" + srcPort[1] + "~" + dstPort[1] + "~" + string(packet.ApplicationLayer().Payload()) + "~\n"
				_, err := file.WriteString(line)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}
			if (srcPort[1] == "22(ssh)" && dstPort[1] == "22(ssh)") || (srcPort[1] == "23(telnet)" && dstPort[1] == "23(telnet)") {
				line := srcIP[1] + "~" + dstIP[1] + "~" + srcPort[1] + "~" + dstPort[1] + "~" + string(packet.ApplicationLayer().Payload()) + "~\n"
				_, err := file.WriteString(line)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}
			fmt.Println(count)
			count += 1
			if count >= 1999 {
				break
			}
		}
	}
	A.MakeReport()
	return
}

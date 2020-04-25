package main

import (
	L "./Pcap"
	"fmt"
	"net"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func listInterfaces() []string {
	interfaces, err := net.Interfaces()
	var ifaceArr []string

	if err != nil {
		fmt.Print(fmt.Errorf("listInterfaces: %+v\n", err.Error()))
		os.Exit(1)
	}

	for _, iface := range interfaces {
		arr,_ := iface.Addrs()
		if len(arr) == 0{
			continue
		}
		ifaceArr = append(ifaceArr, iface.Name)
	 }

	for index, iface := range ifaceArr {
		fmt.Println(index,": ",iface)
	}

	return ifaceArr
}

 func readUserInput(size int) int {
	 //size := len(net.Interfaces())
	 reader := bufio.NewReader(os.Stdin)
	 count := 0

	 for {
		 fmt.Print("Enter index of the device: ")
		 indexStr, _ := reader.ReadString('\n')
		 indexStr = strings.Replace(indexStr, "\n", "", -1)
		 index, convErr := strconv.Atoi(indexStr)

		 if convErr != nil {
			 fmt.Println("Cannot convert input to int")
			 os.Exit(1)
		 }

		 if index >= 0 && index < size {
			 return index
		}

		if count == 2 {
			fmt.Println("Invalid choices made. Shutting down...")
			os.Exit(1)
		}

		fmt.Println("Invalid choice. ( 0 -",size-1,")")
	 }
 }

 func main() {
	 var iface []string
	 fmt.Println("|:|:|:|:|:| GoShark |:|:|:|:|:|")
	 fmt.Println("Available connected interfaces")

	 iface = listInterfaces()
	 input := readUserInput(len(iface))
	 fmt.Println(input)
	 L.Capture(iface[input])
	 // list off interfaces
	 // have user choose 1 AND/OR 2,....
	 // then call capture.go with that interface
 }

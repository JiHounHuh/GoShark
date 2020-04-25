package main

import (
	L "./Pcap"
	"fmt"
	"net"
	"bufio"
	"os"
	"strconv"
 )

 func listInterfaces() {
	 interfaces, err := net.Interfaces()
	 addrs, addrErr := net.InterfaceAddrs()

	 if err != nil {
		 fmt.Print(fmt.Errorf("listInterfaces: %+v\n", err.Error()))
		 return
	 }

	 if addrErr != nil {
		fmt.Println(fmt.Errorf("Addr: %+v\n", addrErr.Error()))
		return
	 }

	 for index, iface := range interfaces {
		 toCheck := addrs[index].String()

		 if toCheck == "::1/128" { // make regex to match other subnets
			 toCheck = "Not Connected"
		 }

		 if toCheck == "0.0.0.0/24" {
			 toCheck = "Not Connected"
		 }

		 fmt.Println(index,": ",iface.Name, " : ", toCheck)
	 }
 }

 func readUserInput(size int) int {
	 //size := len(net.Interfaces())
	 reader := bufio.NewReader(os.Stdin)
	 count := 0

	 for {
		 fmt.Print("Enter index of the device: ")
  	 indexStr, _ := reader.ReadString('\n')
  	 index, _ := strconv.Atoi(indexStr)

  	 if index > 0 && index < size {
  		 return index
  	 }
		 if count == 2 {
			 fmt.Println("Invalid choices made. Shutting down...")
			 os.Exit(1)
		 }

		 fmt.Println("Invalid choice. (0-",size-1,")")
	 }
 }

 func main() {
	 fmt.Println("|:|:|:|:|:| GoShark |:|:|:|:|:|")
	 fmt.Println("Available interfaces")
	 iface, _ := net.Interfaces()
	 listInterfaces()
	 input := readUserInput(len(iface))
	 fmt.Println(input)
	 // list off interfaces
	 // have user choose 1 AND/OR 2,....
	 // then call capture.go with that interface
 }

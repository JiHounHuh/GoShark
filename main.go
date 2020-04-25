package main

import (
	"fmt"
	"net"
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

 func main() {
	 fmt.Println("|:|:|:|:|:| GoShark |:|:|:|:|:|")
	 fmt.Println("Available interfaces")
	 listInterfaces()
	 // list off interfaces
	 // have user choose 1 AND/OR 2,....
	 // then call capture.go with that interface
 }

package main

import (
	"fmt"
	"net"
 )

 func listInterfaces() {
	 interfaces, err := net.Interfaces()
	 if err != nil {
		 fmt.Print(fmt.Errorf("listInterfaces: %+v\n", err.Error()))
		 return
	 }
	 for index, iface := range interfaces {
		 fmt.Println(index,": ",iface.Name)
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

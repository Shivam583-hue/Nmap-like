package main

import (
	"Nmap/CheckOs"
	"Nmap/FindIp"
	"Nmap/identifyServiceOnPort"
	"flag"
	"fmt"
)

func main() {
	port := flag.String("port", "", "port to scan")
	os := flag.Bool("os", false, "check os")
	ip := flag.Bool("ip", false, "find ip")
	flag.Parse()

	if *os {
		CheckOs.CheckOs()
		return
	}

	if *ip {
		FindIp.FindIp()
		return
	}

	host := "localhost"

	service, err := identifyServiceOnPort.IdentifyServiceOnPort(host, *port)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(service)
	}
}

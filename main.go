package main

import (
	"Nmap/CheckHTTPMethods"
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
	url := flag.String("url", "", "url to scan")
	flag.Parse()

	if *os {
		CheckOs.CheckOs()
		return
	}

	if *url != "" {
		CheckHTTPMethods.CheckHTTPMethods(*url)
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

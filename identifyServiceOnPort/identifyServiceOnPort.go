package identifyServiceOnPort

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"time"
)

func IdentifyServiceOnPort(host string, port string) (string, error) {
	portNum, err := strconv.Atoi(port)
	if err != nil {
		return "", fmt.Errorf("invalid port: %v", err)
	}

	address := fmt.Sprintf("%s:%d", host, portNum)
	conn, err := net.DialTimeout("tcp", address, 3*time.Second)
	if err != nil {
		return "", fmt.Errorf("unable to connect to port %d: %v", portNum, err)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost: %s\r\n\r\n", host)

	conn.SetReadDeadline(time.Now().Add(3 * time.Second))
	reader := bufio.NewReader(conn)
	banner, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("error reading banner: %v", err)
	}

	return fmt.Sprintf("Service running on port %d: %s", portNum, banner), nil
}

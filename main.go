package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	// set the env var for testing
	os.Setenv("TCPHEALTH_HOST_01", "www.google.com")
	os.Setenv("TCPHEALTH_PORT_01", "80")

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if strings.Contains(pair[0], "TCPHEALTH_HOST_") {
			//fmt.Println(pair[0] + "=" + pair[1])
			ID := strings.SplitN(pair[0], "_", 3)
			PORT := os.Getenv("TCPHEALTH_PORT_" + ID[2])
			fmt.Println("-- Attempting connection to " + pair[1] + "on " + PORT + " --")
			raw_connect(pair[1], PORT)
		}
	}
}

func raw_connect(host string, port string) {
	timeout := time.Second
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		fmt.Println("Connecting error:", err)
	}
	if conn != nil {
		defer conn.Close()
		fmt.Println("Successfully opened TCP connection", net.JoinHostPort(host, port))
	}
}

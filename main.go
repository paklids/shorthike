package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	// set the env var
	os.Setenv("TCPHEALTH_01_HOST_", "www.google.com")
	// read the env var
	//fmt.Println("TCPHEALTH_01_HOST_:", os.Getenv("TCPHEALTH_01_HOST_"))

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if strings.Contains(pair[0], "TCPHEALTH_") {
			fmt.Println(pair[0] + "=" + pair[1])
			raw_connect("8.8.8.8", "53")
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
		fmt.Println("Opened", net.JoinHostPort(host, port))
	}
}

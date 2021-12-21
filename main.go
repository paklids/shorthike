package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

//connect_timeout represented like 2s (2 seconds) or like 30ms (30 milliseconds)
var connect_timeout, _ = time.ParseDuration(getEnv("TCPHEALTH_CONNECT_TIMEOUT", "1s"))

//run_interval can be represented the same way - but never shorter than connect_interval
var run_interval, _ = time.ParseDuration(getEnv("TCPHEALTH_RUN_INTERVAL", "10s"))

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	if connect_timeout > run_interval {
		fmt.Fprintf(os.Stderr, "Connection timeout is too high\n")
		os.Exit(1)
	}
	dispatcherTicker := time.NewTicker(time.Duration(run_interval))
	for {
		select {
		case <-dispatcherTicker.C:
			dispatcher()
		}
	}
}

func dispatcher() {
	// set the env var for testing
	//os.Setenv("TCPHEALTH_HOST_01", "www.google.com")
	//os.Setenv("TCPHEALTH_PORT_01", "80")

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if strings.Contains(pair[0], "TCPHEALTH_HOST_") {
			//fmt.Println(pair[0] + "=" + pair[1])
			ID := strings.SplitN(pair[0], "_", 3)
			PORT := os.Getenv("TCPHEALTH_PORT_" + ID[2])
			addr := checkIPAddress(pair[1])
			//fmt.Println("-- Attempting connection to ", pair[1], " on ", PORT, " --")
			raw_connect(pair[1], addr, PORT)
		}
	}
}

func checkIPAddress(ip string) (addr string) {
	if net.ParseIP(ip) == nil {
		//fmt.Printf("Host: %s - is NOT an IP address\n", ip)
		addr, err := net.LookupIP(ip)
		if err != nil {
			fmt.Println("Host", ip, "does not resolve properly")
		} else {
			//fmt.Println("Host "+ip+" resolves to: ", addr)
			return addr[0].String()
		}
	} else {
		fmt.Printf("Host: %s - is an IP address\n", ip)
		return ip
	}
	return
}

func raw_connect(host string, addr string, port string) {
	fmt.Println("Connection timeout is ", connect_timeout)
	timeout := time.Duration(connect_timeout)
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(addr, port), timeout)
	if err != nil {
		fmt.Println("Failure connecting to ", host, " - ", net.JoinHostPort(addr, port), "within", timeout, err)
	}
	if conn != nil {
		defer conn.Close()
		fmt.Println("Successfully opened TCP connection to ", host, " - ", net.JoinHostPort(addr, port))
	}
}

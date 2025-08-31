package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
)

type NetMap struct {
	Src      int    `json:"src"`
	Dst      int    `json:"dst"`
	EndPoint string `json:"endpoint"`
}

type Config struct {
	Mappings []NetMap `json:"mappings"`
}

var MAPPER []NetMap

func simpleForward(src, dst net.Conn) {
	defer src.Close()
	defer dst.Close()

	_, err := io.Copy(dst, src)
	if err != nil {
		fmt.Println("Error copying data:", err)
		return
	}
}

func handleConn(c net.Conn, dst int, endpoint string) {
	remote, err := net.Dial("tcp", fmt.Sprintf("%s:%d", endpoint, dst))
	if err != nil {
		fmt.Println("Error dialing remote:", err)
		c.Close()
		return
	}
	go simpleForward(c, remote)
	go simpleForward(remote, c)
}

func proxy(c net.Listener, dst int, endpoint string) {
	for {
		conn, err := c.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn, dst, endpoint)
	}
}

func loadConfig(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read config file %s: %v", filename, err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return fmt.Errorf("failed to parse config file: %v", err)
	}

	MAPPER = config.Mappings
	return nil
}

func main() {
	configFile := "/etc/labproxy.json"
	if len(os.Args) > 1 {
		configFile = os.Args[1]
	}

	if err := loadConfig(configFile); err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	if len(MAPPER) == 0 {
		fmt.Println("No mappings found in config file")
		os.Exit(1)
	}

	for _, m := range MAPPER {
		fmt.Println("Mapping ", m.Src, " to ", m.Dst, " on ", m.EndPoint)
		l, e := net.Listen("tcp", fmt.Sprintf(":%d", m.Src))
		if e != nil {
			panic(e)
		}
		go proxy(l, m.Dst, m.EndPoint)
	}

	select {}
}

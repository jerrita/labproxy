package main

import (
	"flag"
	"fmt"
	"io"
	"net"
)

type NetMap struct {
	Src int
	Dst int
}

var Config struct {
	EndPoint string
	Api      string
	Password string
}

var MAPPER = [...]NetMap{
	{80, 80},
	{443, 443},
}

func simpleForward(src, dst net.Conn) {
	defer src.Close()
	defer dst.Close()

	_, err := io.Copy(dst, src)
	if err != nil {
		fmt.Println("Error copying data:", err)
		return
	}
}

func handleConn(c net.Conn, dst int) {
	remote, err := net.Dial("tcp", fmt.Sprintf("%s:%d", Config.EndPoint, dst))
	if err != nil {
		fmt.Println("Error dialing remote:", err)
		c.Close()
		return
	}
	go simpleForward(c, remote)
	go simpleForward(remote, c)
}

func proxy(c net.Listener, dst int) {
	for {
		conn, err := c.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn, dst)
	}
}

func main() {
	ep := flag.String("endpoint", "localhost", "The endpoint to proxy to (domain/ip)")
	api := flag.String("api", ":4136", "Manager api path")
	password := flag.String("password", "", "Manager password")
	flag.Parse()

	Config.EndPoint = *ep
	Config.Api = *api
	Config.Password = *password

	fmt.Println("Manager API on", Config.Api)

	for _, m := range MAPPER {
		fmt.Println("Mapping ", m.Src, " to ", m.Dst)
		l, e := net.Listen("tcp", fmt.Sprintf(":%d", m.Src))
		if e != nil {
			panic(e)
		}
		go proxy(l, m.Dst)
	}

	select {}
}

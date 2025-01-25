package main

import (
	"flag"
	"fmt"
	"io"
	"net"
)

var Config struct {
	EndPoint string
	Api      string
	Password string
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

func handleConn(c net.Conn) {
	remote, err := net.Dial("tcp", Config.EndPoint)
	if err != nil {
		fmt.Println("Error dialing remote:", err)
		c.Close()
		return
	}
	go simpleForward(c, remote)
	go simpleForward(remote, c)
}

func proxy(c net.Listener) {
	for {
		conn, err := c.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn)
	}
}

func main() {
	ep := flag.String("endpoint", "localhost:8080", "The endpoint to proxy to (tcp only)")
	api := flag.String("api", ":4136", "Manager api path")
	password := flag.String("password", "", "Manager password")
	flag.Parse()

	Config.EndPoint = *ep
	Config.Api = *api
	Config.Password = *password

	fmt.Println("Starting proxy on :80 and :443, forwarding to", Config.EndPoint)
	fmt.Println("Manager API on", Config.Api)

	l, e := net.Listen("tcp", ":80")
	if e != nil {
		panic(e)
	}
	go proxy(l)
	l, e = net.Listen("tcp", ":443")
	if e != nil {
		panic(e)
	}
	proxy(l)
}

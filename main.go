package main

import (
	"Share/client"
	"Share/server"
	"flag"
	"fmt"
	"net"
	_ "net/http/pprof"
)

func main() {
	var filePath string
	var Port string

	flag.StringVar(&filePath, "PATH", "//", "abc/d/e")
	flag.StringVar(&Port, "PORT", ":4444", "1234")
	flag.Parse()
	if filePath != "//" {
		addr, err := net.ResolveUDPAddr("udp4", Port)
		if err != nil {
			panic(err)
		}
		server.Send(filePath, addr)
	} else {
		addr, err := net.ResolveUDPAddr("udp4", Port)
		if err != nil {
			panic(err)
		}
		fmt.Println("addr of server:", addr)
		client.Recv(addr)
	}

}

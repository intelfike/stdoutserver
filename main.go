package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

var port = flag.String("port", ":8888", "Listen port number.")

func init() {
	flag.Parse()
}

func main() {
	ln, err := net.Listen("tcp", *port)
	if err != nil {
		panic(err.Error())
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	io.Copy(os.Stdout, conn)
}

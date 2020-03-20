package main

import (
	"log"
	"net"
	"io"
)


func handle(src net.Conn) {
	dst, err := net.Dial("tcp", "rebay1982.github.io:80")
	if err != nil {
		log.Fatalln("Could not connect to destination")
	}
	defer dst.Close()

	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln("Could'nt forward to destination")
		}
	}()

	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}


func main() {
	// Listen for incoming connections
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalln("Couldn't listen on port 8080")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connections")
		}
		go handle(conn)
	}
}

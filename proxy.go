package main

import (
	"log"
	"net"
	"io"
)

func echo(conn net.Conn) {
	defer conn.Close()

	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("Failed to read/write")
	}

	// #1
	// Buffer to store received data.
	//b := make([]byte, 512)

	// #1, 2
	//for {
		/*
		// #2
		reader := bufio.NewReader(conn)
		s, err := reader.ReadString('\n')

		if err != nil {
			log.Fatalln("Failed to read string")
		}

		log.Printf("Read [%d] bytes: %s", len(s), s)

		writer := bufio.NewWriter(conn)
		_, err = writer.WriteString(s)

		if err != nil {
			log.Fatalln("Failed to write string")
		}

		writer.Flush()
		*/
		/*
		// #1
		// Receive data via conn.
		size, err := conn.Read(b[0:])

		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}

		if err != nil {
			log.Println("Unexpected error")
			break
		}

		log.Printf("Received [%d] bytes: %s\n", size, string(b))

		// Send data via conn.Write
		log.Println("Writing Data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}
		*/
	//}
}

func main() {
	// Bind to TCP port 9292 on all interfaces.
	listener, err := net.Listen("tcp", ":9292")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}

	log.Println("Listening on 0.0.0.0:9292")
	for {
		conn, err := listener.Accept()
		log.Println("Receive connection")
		if err != nil {
			log.Fatalln("Unable to receive connection")
		}

		go echo(conn)
	}
}

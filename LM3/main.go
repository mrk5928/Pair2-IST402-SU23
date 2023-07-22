package main

import (
	"bufio"
	"crypto/tls"
	"log"
	"net"
)

func main() {
	// Set log flags to include the filename and line number in log messages
	log.SetFlags(log.Lshortfile)

	// Load the X.509 certificate and private key for TLS encryption
	cer, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Println(err)
		return
	}

	// Create a TLS configuration using the loaded certificate and key
	config := &tls.Config{Certificates: []tls.Certificate{cer}}

	// Start a TLS listener on port 443 with the TLS configuration
	ln, err := tls.Listen("tcp", ":443", config)
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()

	// Accept and handle incoming connections in a loop
	for {
		// Wait for a new client connection
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		// Handle the connection in a separate goroutine to allow concurrent connections
		go handleConnection(conn)
	}
}

// handleConnection handles an individual client connection
func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Create a buffered reader to read data from the connection
	r := bufio.NewReader(conn)

	// Continuously read incoming messages from the client
	for {
		// Read the message until a newline character is encountered
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}

		// Print the received message to the server's console
		println(msg)

		// Write "world" back to the client
		n, err := conn.Write([]byte("world\n"))
		if err != nil {
			log.Println(n, err)
			return
		}
	}
}

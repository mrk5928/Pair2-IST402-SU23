package main

import (
	"crypto/tls"
	"log"
)

func main() {
	// Set log flags to include the filename and line number in the log output
	log.SetFlags(log.Lshortfile)

	// Create a TLS configuration with InsecureSkipVerify set to true to skip certificate verification
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	// Dial a TLS connection to the server at 127.0.0.1:443
	conn, err := tls.Dial("tcp", "127.0.0.1:443", conf)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close() // Ensure the connection is closed when the function returns or exits

	// Write "hello\n" to the TLS connection (server)
	n, err := conn.Write([]byte("hello\n"))
	if err != nil {
		log.Println(n, err)
		return
	}

	// Read the server's response from the connection into a buffer (buf)
	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		log.Println(n, err)
		return
	}

	// Print the server's response as a string
	println(string(buf[:n]))
}

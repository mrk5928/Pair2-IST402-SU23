// LM4: ChaCha20 Programming
package main

import (
	//"bufio"
	"crypto/rand"
	//"encoding/hex"
	"fmt"
	"io"
	//"os"
	//"golang.org/x/crypto/chacha20"
)

func main() {
	// Generate a random 256-bit key
	key := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		fmt.Println("Failed to generate random key:", err)
		return
	}

}

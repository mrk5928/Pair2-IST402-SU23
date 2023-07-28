// LM4: ChaCha20 Programming
package main

import (
	//"bufio"
	"crypto/rand"
	"fmt"
	//"golang.org/x/crypto/chacha20"
	"io"
	//"os"
)

func main() {
	// Generate a random 256-bit key
	key := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		fmt.Println("Failed to generate random key:", err)
		return
	}
	// Get user input for plaintext
	//var plaintext string
	//fmt.Println("Enter plaintext: ")
	//scanner := bufio.NewScanner(os.Stdin)
	//if scanner.Scan() {
	//plaintext = scanner.Text()
	//}

	// Convert plaintext to byte slice
	//plaintextBytes := []byte(plaintext)

	// Generate a random 96-bit nonce
	//nonce := make([]byte, chacha20.NonceSizeX)
	//if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
	//fmt.Println("Failed to generate random nonce:", err)
	//return
	//}

	// Create a new ChaCha20 cipher with the random key and nonce
	//c, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	//if err != nil {
	//fmt.Println("Failed to create ChaCha20 cipher:", err)
	//return
	//}
}

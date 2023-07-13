package main

import (
	"fmt"
	"strconv"
)

var codebook = [4][2]int{{0b00, 0b01}, {0b01, 0b10}, {0b10, 0b11}, {0b11, 0b00}} // Codebook used for encryption and decryption

var iv int = 0b10 // Initialization Vector (IV) for OFB mode

// Function to look up the codebook and find the corresponding ciphertext value for a given XOR value
func codebookLookup(xor int) int {
	for i := 0; i < 4; i++ {
		if codebook[i][0] == xor {
			return codebook[i][1]
		}
	}
	return 0
}

// Function to reverse lookup the codebook and find the corresponding plaintext value for a given ciphertext
func codebookReverseLookup(ciphertext int) int {
	for i := 0; i < 4; i++ {
		if codebook[i][1] == ciphertext {
			return codebook[i][0]
		}
	}
	return 0
}

func main() {
	var plaintextStr string
	fmt.Print("Enter the plaintext (a sequence of 2-bit values): ")
	fmt.Scan(&plaintextStr)

	plaintextLen := len(plaintextStr)
	message := make([]int, plaintextLen)

	for i, char := range plaintextStr {
		bit, err := strconv.Atoi(string(char))
		if err != nil || (bit != 0 && bit != 1) {
			fmt.Println("Invalid input. The plaintext should only consist of 0s and 1s.")
			return
		}
		message[i] = bit
	}

	fmt.Println("ECB encryption details:")
	fmt.Printf("Plaintext: %b\n", message)
	ciphertext := make([]int, plaintextLen)

	// Encryption using ECB mode
	for i := 0; i < plaintextLen; i++ {
		ciphertext[i] = codebookLookup(message[i])
		fmt.Printf("The ciphered value of %b is %b\n", message[i], ciphertext[i])
	}

	fmt.Println("\nECB decryption details:")
	decryptedPlaintext := make([]int, plaintextLen)

	// Decryption using ECB mode
	for i := 0; i < plaintextLen; i++ {
		decryptedPlaintext[i] = codebookReverseLookup(ciphertext[i])
		fmt.Printf("The deciphered value of %b is %b\n", ciphertext[i], decryptedPlaintext[i])
	}

	fmt.Printf("\nDecrypted ECB message: %b\n", decryptedPlaintext)

	fmt.Println("\nOFB encryption details:")
	fmt.Printf("Plaintext: %b\n", message)
	stream := iv
	ciphertext = make([]int, plaintextLen)

	// Encryption using OFB mode
	for i := 0; i < plaintextLen; i++ {
		xor := message[i] ^ stream
		ciphertext[i] = codebookLookup(xor)
		stream = codebookLookup(stream) // Update the stream value using codebook lookup
		fmt.Printf("The ciphered value of %b is %b\n", message[i], ciphertext[i])
	}

	fmt.Println("\nOFB decryption details:")
	stream = iv
	decryptedPlaintext = make([]int, plaintextLen)

	// Decryption using OFB mode
	for i := 0; i < plaintextLen; i++ {
		xor := ciphertext[i] ^ stream
		decryptedPlaintext[i] = codebookReverseLookup(xor)
		stream = codebookLookup(ciphertext[i]) // Update the stream value using codebook lookup
		fmt.Printf("The deciphered value of %b is %b\n", ciphertext[i], decryptedPlaintext[i])
	}

	fmt.Printf("\nDecrypted OFB message: %b\n", decryptedPlaintext)
}

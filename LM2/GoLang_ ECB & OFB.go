package main

import "fmt"

// Codebook (substitute with your own block cipher implementation if needed)
var Codebook = [4][2]int{{0b00, 0b01}, {0b01, 0b10}, {0b10, 0b11}, {0b11, 0b00}}

// ECB encryption
func ecbEncrypt(plaintext []int) []int {
	ciphertext := make([]int, len(plaintext))

	for i := 0; i < len(plaintext); i++ {
		ciphertext[i] = codebookLookup(plaintext[i])
	}

	return ciphertext
}

// ECB decryption
func ecbDecrypt(ciphertext []int) []int {
	plaintext := make([]int, len(ciphertext))

	for i := 0; i < len(ciphertext); i++ {
		plaintext[i] = codebookLookupbyValue(ciphertext[i])
	}

	return plaintext
}

// OFB encryption
func ofbEncrypt(plaintext []int, iv int) []int {
	ciphertext := make([]int, len(plaintext))
	keystream := generateKeystream(iv, len(plaintext))

	for i := 0; i < len(plaintext); i++ {
		ciphertext[i] = keystream[i] ^ plaintext[i]
	}

	return ciphertext
}

// OFB decryption
func ofbDecrypt(ciphertext []int, iv int) []int {
	plaintext := make([]int, len(ciphertext))
	keystream := generateKeystream(iv, len(ciphertext))

	for i := 0; i < len(ciphertext); i++ {
		plaintext[i] = keystream[i] ^ ciphertext[i]
	}

	return plaintext
}

func codebookLookup(xor int) int {
	for i := 0; i < 4; i++ {
		if Codebook[i][0] == xor {
			return Codebook[i][1]
		}
	}
	return 0
}

func codebookLookupbyValue(xor int) int {
	for i := 0; i < 4; i++ {
		if Codebook[i][1] == xor {
			return Codebook[i][0]
		}
	}
	return 0
}

func generateKeystream(iv int, length int) []int {
	keystream := make([]int, length)
	keystream[0] = codebookLookup(iv) // Use IV to start the keystream

	for i := 1; i < length; i++ {
		keystream[i] = codebookLookup(keystream[i-1]) // Generate the next keystream block
	}

	return keystream
}

var message = [4]int{0b01, 0b00, 0b10, 0b00}
var IV = 0b10 // Initialization Vector, it should be random and unique for each encryption

func main() {
	// Convert the message array to a slice
	plaintext := message[:]

	// ECB encryption and decryption
	fmt.Println("ECB Encryption and Decryption:")
	ciphertextECB := ecbEncrypt(plaintext)
	fmt.Println("The plaintext values:")
	for i := 0; i < len(plaintext); i++ {
		fmt.Printf("The plaintext value of a is %02b\n", plaintext[i])
	}

	fmt.Println("The ciphered values:")
	for i := 0; i < len(ciphertextECB); i++ {
		fmt.Printf("The ciphered value of a is %02b\n", ciphertextECB[i])
	}

	decryptedECB := ecbDecrypt(ciphertextECB)
	fmt.Println("The decrypted plaintext values:")
	for i := 0; i < len(decryptedECB); i++ {
		fmt.Printf("The plaintext value of a is %02b\n", decryptedECB[i])
	}

	fmt.Println()

	// OFB encryption and decryption
	fmt.Println("OFB Encryption and Decryption:")
	ciphertextOFB := ofbEncrypt(plaintext, IV)
	fmt.Println("The plaintext values:")
	for i := 0; i < len(plaintext); i++ {
		fmt.Printf("The plaintext value of a is %02b\n", plaintext[i])
	}

	fmt.Println("The ciphered values:")
	for i := 0; i < len(ciphertextOFB); i++ {
		fmt.Printf("The ciphered value of a is %02b\n", ciphertextOFB[i])
	}

	decryptedOFB := ofbDecrypt(ciphertextOFB, IV)
	fmt.Println("The decrypted plaintext values:")
	for i := 0; i < len(decryptedOFB); i++ {
		fmt.Printf("The plaintext value of a is %02b\n", decryptedOFB[i])
	}
}

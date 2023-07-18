package main

import "fmt"

/* an array with 4 rows and 2 columns */
var codebook = [4][2]int{{0b00, 0b01}, {0b01, 0b10}, {0b10, 0b11}, {0b11, 0b00}}
var messages = [4]int{0b01, 0b00, 0b10, 0b00}
var cipherCBC = [4]int{}
var cipherCFB = [4]int{}
var iv int = 0b10

func codesbookLookup(xor int) int {
	var i, j int = 0, 0
	for i = 0; i < 4; i++ {
		if codebook[i][j] == xor {
			j++
			return codebook[i][j]
		}
	}
	return -1 // Return -1 if the lookup fails
}

func encryptBlockCBC(plaintext, iv int) int {
	xor := plaintext ^ iv
	return codesbookLookup(xor)
}

func decryptBlockCBC(ciphertext, iv int) int {
	xor := codesbookLookup(ciphertext) ^ iv
	return xor
}

func encryptBlockCFB(plaintext, iv int) int {
	xor := plaintext ^ iv
	return xor
}

func decryptBlockCFB(ciphertext, iv int) int {
	xor := ciphertext ^ iv
	return xor
}

func main() {
	// CBC Mode
	var xor int = 0
	var lookupValue int = 0
	lookupValue = codesbookLookup(iv)

	// Display the original message
	fmt.Println("CBC Mode:")
	for i := 0; i < 4; i++ {
		fmt.Printf("The plaintext value of a is %02b\n", messages[i])
	}

	// Encryption (CBC)
	for i := 0; i < 4; i++ {
		xor = messages[i] ^ lookupValue
		lookupValue = codesbookLookup(xor)
		fmt.Printf("The ciphered value of a in CBC mode is %02b\n", xor)
		cipherCBC[i] = xor
	}

	// Decryption (CBC)
	lookupValue = codesbookLookup(iv)
	for i := 0; i < 4; i++ {
		xor = cipherCBC[i] ^ lookupValue
		lookupValue = codesbookLookup(cipherCBC[i])
		fmt.Printf("The plaintext value of a in CBC mode is %02b\n", xor)
	}

	fmt.Println()

	// CFB Mode
	lookupValue = codesbookLookup(iv)

	// Display the original message
	fmt.Println("CFB Mode:")
	for i := 0; i < 4; i++ {
		fmt.Printf("The plaintext value of a is %02b\n", messages[i])
	}

	// Encryption (CFB)
	for i := 0; i < 4; i++ {
		cipherCFB[i] = encryptBlockCFB(messages[i], lookupValue)
		lookupValue = cipherCFB[i]
		fmt.Printf("The ciphered value of a in CFB mode is %02b\n", cipherCFB[i])
	}

	// Decryption (CFB)
	lookupValue = codesbookLookup(iv)
	for i := 0; i < 4; i++ {
		plaintext := decryptBlockCFB(cipherCFB[i], lookupValue)
		lookupValue = cipherCFB[i]
		fmt.Printf("The plaintext value of a in CFB mode is %02b\n", plaintext)
	}
}

package main

import (
	"fmt"
	"strconv"
)

var codeBook = [][]int{
	{0b00, 0b01},
	{0b01, 0b10},
	{0b10, 0b11},
	{0b11, 0b00},
}

// Function to perform a lookup in the code book and return the corresponding value
func codebook_Lookup(xor int) int {
	for i := 0; i < len(codeBook); i++ {
		if codeBook[i][0] == xor {
			return codeBook[i][1]
		}
	}
	return 0
}

func main() {
	var plaintext string
	fmt.Print("Enter the plaintext (binary): ")
	fmt.Scanln(&plaintext)

	var message []int
	for _, bit := range plaintext {
		num, _ := strconv.Atoi(string(bit))
		message = append(message, num)
	}

	var message1 []string
	for _, bit := range message {
		message1 = append(message1, strconv.Itoa(bit))
	}

	fmt.Println("\nCBC encryption details:")
	fmt.Printf("Plaintext: %v\n", message1)

	iv := 0b10
	stream := iv
	var ciphertext []string
	for _, bit := range message {
		xor := bit ^ stream
		ciphertext = append(ciphertext, strconv.Itoa(codebook_Lookup(xor)))
		stream = codebook_Lookup(xor)
		fmt.Printf("The ciphered value of %04b is %04b\n", bit, codebook_Lookup(xor))
	}

	reverseSlice(ciphertext)
	reverseSlice(message1)

	fmt.Println("\nCBC decryption details:")
	fmt.Printf("Ciphertext: %v\n", ciphertext)

	stream = iv
	var decryptedMessage []int
	for _, bit := range ciphertext {
		cipherInt, _ := strconv.Atoi(bit)
		xor := codebook_Lookup(cipherInt)
		decryptedMessage = append(decryptedMessage, xor^stream)
		stream = cipherInt
		fmt.Printf("The deciphered value of %04b is %04b\n", cipherInt, xor^stream)
	}

	reverseSliceInt(decryptedMessage)
	var originalPlaintext string
	for _, bit := range decryptedMessage {
		originalPlaintext += strconv.Itoa(bit)
	}
	fmt.Printf("\nOriginal message: %v\n", originalPlaintext)

	fmt.Println("\nCFB encryption details:")
	fmt.Printf("Plaintext: %v\n", message1)

	stream = iv
	ciphertext = nil
	for _, bit := range message {
		xor := bit ^ stream
		ciphertext = append(ciphertext, strconv.Itoa(codebook_Lookup(xor)))
		stream = codebook_Lookup(xor) ^ iv
		fmt.Printf("The ciphered value of %04b is %04b\n", bit, codebook_Lookup(xor))
	}

	reverseSlice(ciphertext)
	reverseSlice(message1)

	fmt.Println("\nCFB decryption details:")
	fmt.Printf("Ciphertext: %v\n", ciphertext)

	stream = iv
	decryptedMessage = nil
	for _, bit := range ciphertext {
		cipherInt, _ := strconv.Atoi(bit)
		xor := codebook_Lookup(cipherInt)
		decryptedMessage = append(decryptedMessage, xor^stream)
		stream = cipherInt ^ decryptedMessage[len(decryptedMessage)-1]
		fmt.Printf("The deciphered value of %04b is %04b\n", cipherInt, xor^stream)
	}

	reverseSliceInt(decryptedMessage)
	originalPlaintext = ""
	for _, bit := range decryptedMessage {
		originalPlaintext += strconv.Itoa(bit)
	}
	fmt.Printf("\nOriginal message: %v\n", originalPlaintext)
}

// Function to reverse the order of elements in a string slice
func reverseSlice(slice []string) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Function to reverse the order of elements in an int slice
func reverseSliceInt(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

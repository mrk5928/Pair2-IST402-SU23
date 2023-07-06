package main

import (
	"fmt"
	"strings"
)

func decrypt(n int, ciphertext string) string {
	result := ""
	for _, l := range ciphertext {
		if strings.ContainsRune("ABCDEFGHIJKLMNOPQRSTUVWXYZ", l) {
			index := strings.IndexRune("ABCDEFGHIJKLMNOPQRSTUVWXYZ", l)
			i := (index - n + 26) % 26
			result += string("ABCDEFGHIJKLMNOPQRSTUVWXYZ"[i])
		} else if strings.ContainsRune("abcdefghijklmnopqrstuvwxyz", l) {
			index := strings.IndexRune("abcdefghijklmnopqrstuvwxyz", l)
			i := (index - n + 26) % 26
			result += string("abcdefghijklmnopqrstuvwxyz"[i])
		} else {
			result += string(l)
		}
	}
	return result
}

func bruteForce(ciphertext string) {
	for key := 1; key <= 25; key++ {
		plaintext := decrypt(key, ciphertext)
		fmt.Printf("Key: %d\tPlaintext: %s\n", key, plaintext)
	}
}

func main() {
	ciphertext := "Ugew gnwj zwjw Oslkgf"
	bruteForce(ciphertext)
}

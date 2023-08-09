package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Rotor represents a rotor with a substitution mapping and a notch position.
type Rotor struct {
	mapping  string
	notch    rune
	position int
}

func (r *Rotor) setStartPosition(position int) {
	r.position = position
}

func (r *Rotor) substitute(char rune, forward bool) rune {
	offset := r.position
	if !forward {
		offset = -offset
	}
	index := (int(char-'A') + offset + 26) % 26
	return rune(r.mapping[index])
}

func (r *Rotor) rotate() {
	r.position = (r.position + 1) % 26
}

// EnigmaMachine represents an Enigma-like machine with multiple rotors and a reflector.
type EnigmaMachine struct {
	rotors    []*Rotor
	reflector *Rotor
}

func (em *EnigmaMachine) setRotorPositions(positions []int) {
	for i, position := range positions {
		em.rotors[i].setStartPosition(position)
	}
}

func (em *EnigmaMachine) encrypt(plaintext string) string {
	ciphertext := strings.Builder{}
	plaintext = strings.ToUpper(plaintext)
	for _, char := range plaintext {
		if char >= 'A' && char <= 'Z' {
			for _, rotor := range em.rotors {
				char = rotor.substitute(char, true)
			}
			char = em.reflector.substitute(char, true)
			for i := len(em.rotors) - 1; i >= 0; i-- {
				char = em.rotors[i].substitute(char, false)
			}
			ciphertext.WriteRune(char)
			for _, rotor := range em.rotors {
				rotor.rotate()
			}
		}
	}
	return ciphertext.String()
}

func (em *EnigmaMachine) decrypt(ciphertext string) string {
	plaintext := strings.Builder{}
	ciphertext = strings.ToUpper(ciphertext)
	for _, char := range ciphertext {
		if char >= 'A' && char <= 'Z' {
			for _, rotor := range em.rotors {
				char = rotor.substitute(char, true)
			}
			char = em.reflector.substitute(char, true)
			for i := len(em.rotors) - 1; i >= 0; i-- {
				char = em.rotors[i].substitute(char, false)
			}
			plaintext.WriteRune(char)
			for _, rotor := range em.rotors {
				rotor.rotate()
			}
		}
	}
	return plaintext.String()
}

func main() {
	// Example rotor mappings for demonstration purposes
	rotorI := &Rotor{mapping: "EKMFLGDQVZNTOWYHXUSPAIBRCJ", notch: 'Q'}
	rotorII := &Rotor{mapping: "AJDKSIRUXBLHWTMCQGZNPYFVOE", notch: 'E'}
	rotorIII := &Rotor{mapping: "BDFHJLCPRTXVZNYEIWGAKMUSQO", notch: 'V'}

	// Example reflector mapping
	reflectorB := &Rotor{mapping: "YRUHQSLDPXNGOKMIEBFZCWVJAT", notch: 0}

	// Create an Enigma machine with the specified rotors and reflector
	enigma := &EnigmaMachine{
		rotors:    []*Rotor{rotorI, rotorII, rotorIII},
		reflector: reflectorB,
	}

	// Set initial rotor positions
	enigma.setRotorPositions([]int{0, 0, 0})

	// Get user input for plaintext
	var plaintext string
	fmt.Print("Enter the plaintext message: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		plaintext = scanner.Text()
	}

	// Encrypt the message
	ciphertext := enigma.encrypt(plaintext)

	// Decrypt the ciphertext
	decryptedText := enigma.decrypt(ciphertext)

	fmt.Println("Plaintext:", plaintext)
	fmt.Println("Encrypted Text:", ciphertext)
	fmt.Println("Decrypted Text:", decryptedText)
}

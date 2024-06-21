package main

import (
	"fmt"
	"helium/crypto_algs"
)

func 

func main() {
	// Generate a pair of keys
	bits := 2048
	publicKey, privateKey, err := helium.GenerateKeyPairs(bits)
	if err != nil {
		fmt.Println("Error generating keys:", err)
		return
	}
	fmt.Println("Public Key:", *publicKey)
	fmt.Println("Private Key:", *privateKey)

	// Encrypt a message
	message := "Hello, World!"
	encryptedMessage, err := helium.Encrypt(publicKey, &message)
	if err != nil {
		fmt.Println("Error encrypting message:", err)
		return
	}
	fmt.Println("Encrypted Message:", *encryptedMessage)

	// Decrypt the message
	decryptedMessage, err := helium.Decrypt(privateKey, encryptedMessage)
	if err != nil {
		fmt.Println("Error decrypting message:", err)
		return
	}
	fmt.Println("Decrypted Message:", *decryptedMessage)
}

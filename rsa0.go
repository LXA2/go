package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"io/ioutil"
)



func main(){
	fmt.Println("1:encrypt   2:decrypt")
	var choice int
	for {
		_,err := fmt.Scanln(&choice)
		fmt.Println(choice)
		if (err != nil)||((choice!=1)&&(choice!=2)) {
			fmt.Println("try again.",err)
			continue
		}else if choice == 1{
			fmt.Println("encrypt!!!")
			encrypt_inquiry()
		}else {
			fmt.Println("decrypt!!!")
		}
		break
	}
}

func encrypt_inquiry(){
	fmt.Println("1:choose a public key file   2:generate key pairs")
	var choice int
	var public_key_path string
	var initial_message string
	for {
		_,err := fmt.Scanln(&choice)
		fmt.Println(choice)
		if (err != nil)||((choice!=1)&&(choice!=2)) {
			fmt.Println("try again.",err)
			continue
		}else if choice == 1{
			fmt.Println("enter file path(x/x/x.x)")
			for {
				_,err := fmt.Scanln(&public_key_path)
				fmt.Println(public_key_path)
				if (err != nil) {
					fmt.Println("try again.",err)
					continue
				}
				break
			}
			fmt.Println("enter message:")
			for {
				_,err := fmt.Scanln(&initial_message)
				fmt.Println(initial_message)
				if (err != nil) {
					fmt.Println("try again.",err)
					continue
				}
				message := []byte(initial_message)
				publicKey,err :=loadPublicKeyFromString(public_key)
				if err != nil {
					return
				}
				encryptedMessage, err := rsaEncrypt(publicKey, message)
				if err != nil {
					fmt.Println("Error encrypting message:", err)
					break
				}
				fmt.Println("Encrypted message:", encryptedMessage)
				break
			}
		}else {

			fmt.Println("")

		}
		break
	}
}


/*
func main() {
	// 生成RSA密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Error generating RSA key:", err)
		return
	}
	fmt.Println("generated key:",privateKey)

	// 提取公钥
	publicKey := &privateKey.PublicKey

	// 将公钥和私钥保存到文件
	savePEMKey("private_key.pem", privateKey)
	savePublicPEMKey("public_key.pem", publicKey)

	// 要加密的消息
	message := []byte("Hello, RSA encryption!")

	// 加密消息
	encryptedMessage, err := rsaEncrypt(publicKey, message)
	if err != nil {
		fmt.Println("Error encrypting message:", err)
		return
	}
	fmt.Println("Encrypted message:", encryptedMessage)

	// 解密消息
	decryptedMessage, err := rsaDecrypt(privateKey, encryptedMessage)
	if err != nil {
		fmt.Println("Error decrypting message:", err)
		return
	}
	fmt.Println("Decrypted message:", string(decryptedMessage))
}
*/

// 保存私钥到文件
func savePEMKey(filename string, key *rsa.PrivateKey) {
	outFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error saving PEM key:", err)
		return
	}
	defer outFile.Close()

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(key)
	privateKeyPEM := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	pem.Encode(outFile, &privateKeyPEM)
}

// 保存公钥到文件
func savePublicPEMKey(filename string, pubkey *rsa.PublicKey) {
	outFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error saving PEM public key:", err)
		return
	}
	defer outFile.Close()

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		fmt.Println("Error marshalling public key:", err)
		return
	}

	publicKeyPEM := pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	pem.Encode(outFile, &publicKeyPEM)
}

// 使用公钥加密消息
func rsaEncrypt(pub *rsa.PublicKey, msg []byte) ([]byte, error) {
	hash := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, pub, msg, nil)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

// 使用私钥解密消息
func rsaDecrypt(priv *rsa.PrivateKey, ciphertext []byte) ([]byte, error) {
	hash := sha256.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, priv, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}


//load pubic key from a path
func loadPublicKey(path string) (*rsa.PublicKey, error) {
    pubKeyFile, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer pubKeyFile.Close()

    pemBytes, err := ioutil.ReadAll(pubKeyFile)
    if err != nil {
        return nil, err
    }

    block, _ := pem.Decode(pemBytes)
    if block == nil || block.Type != "PUBLIC KEY" {
        return nil, fmt.Errorf("failed to decode PEM block containing public key")
    }

    pub, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
        return nil, err
    }

    rsaPub, ok := pub.(*rsa.PublicKey)
    if !ok {
        return nil, fmt.Errorf("not an RSA public key")
    }

    return rsaPub, nil
}



//load public key from string
func loadPublicKeyFromString(input string) (*rsa.PublicKey, error) {
	pemBytes := []byte(input)
    block, _ := pem.Decode(pemBytes)
    if block == nil || block.Type != "PUBLIC KEY" {
        return nil, fmt.Errorf("failed to decode PEM block containing public key")
    }

    pub, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
        return nil, err
    }

    rsaPub, ok := pub.(*rsa.PublicKey)
    if !ok {
        return nil, fmt.Errorf("not an RSA public key")
    }

    return rsaPub, nil
}
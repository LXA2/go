package main

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/base64"
    "encoding/pem"
    "errors"
    "fmt"
)

// generateKeyPairs 生成指定位数的RSA密钥对，并返回PEM格式的公钥和私钥。
func generateKeyPairs(bits int) (string, string, error) {
    // 生成RSA私钥
    privateKey, err := rsa.GenerateKey(rand.Reader, bits)
    if err != nil {
        return "", "", err
    }

    // 编码私钥为PEM格式
    privateKeyPEM := pem.EncodeToMemory(&pem.Block{
        Type:  "RSA PRIVATE KEY",
        Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
    })

    // 提取公钥并编码为PEM格式
    publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
    if err != nil {
        return "", "", err
    }
    publicKeyPEM := pem.EncodeToMemory(&pem.Block{
        Type:  "RSA PUBLIC KEY",
        Bytes: publicKeyBytes,
    })

    return string(publicKeyPEM), string(privateKeyPEM), nil
}

// encrypt 使用给定的PEM公钥加密消息，并返回加密后的内容和错误信息。
func encrypt(publicKeyPEM string, message string) (string, error) {
    // 解析PEM编码的公钥
    block, _ := pem.Decode([]byte(publicKeyPEM))
    if block == nil || block.Type != "RSA PUBLIC KEY" {
        return "", errors.New("failed to decode PEM block containing public key")
    }

    pub, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
        return "", err
    }

    rsaPub, ok := pub.(*rsa.PublicKey)
    if !ok {
        return "", errors.New("not an RSA public key")
    }

    // 使用公钥加密消息
    encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPub, []byte(message))
    if err != nil {
        return "", err
    }

    // 返回Base64编码的加密消息
    return base64.StdEncoding.EncodeToString(encryptedBytes), nil
}

// decrypt 使用给定的PEM私钥解密消息，并返回解密后的内容和错误信息。
func decrypt(privateKeyPEM string, encryptedMessage string) (string, error) {
    // 解析PEM编码的私钥
    block, _ := pem.Decode([]byte(privateKeyPEM))
    if block == nil || block.Type != "RSA PRIVATE KEY" {
        return "", errors.New("failed to decode PEM block containing private key")
    }

    priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
        return "", err
    }

    // 解码Base64编码的加密消息
    encryptedBytes, err := base64.StdEncoding.DecodeString(encryptedMessage)
    if err != nil {
        return "", err
    }

    // 使用私钥解密消息
    decryptedBytes, err := rsa.DecryptPKCS1v15(rand.Reader, priv, encryptedBytes)
    if err != nil {
        return "", err
    }

    return string(decryptedBytes), nil
}


func main() {
    // 生成密钥对
    publicKey, privateKey, err := generateKeyPairs(2048)
    if err != nil {
        fmt.Println("Error generating keys:", err)
        return
    }

    fmt.Println("Public Key:\n", publicKey)
	fmt.Println("\n------------------------------------------\n")
    fmt.Println("Private Key:\n", privateKey)
	fmt.Println("\n------------------------------------------\n")


	message := "test message!!! 一二三四五六七八九十"
    encryptedMessage, err := encrypt(publicKey, message)
    if err != nil {
        fmt.Println("Error encrypting message:", err)
        return
    }
    fmt.Println("Encrypted Message:", encryptedMessage)

    // 解密消息
    decryptedMessage, err := decrypt(privateKey, encryptedMessage)
    if err != nil {
        fmt.Println("Error decrypting message:", err)
        return
    }
    fmt.Println("Decrypted Message:", decryptedMessage)
}

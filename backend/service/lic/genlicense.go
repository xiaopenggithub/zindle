package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// 生成RSA密钥对
func generateRSAKeyPair(bits int) (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

// 保存私钥到文件
func savePrivateKey(privateKey *rsa.PrivateKey, filename string) error {
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	err = pem.Encode(file, privateKeyBlock)
	if err != nil {
		file.Close()
		return err
	}
	file.Close()
	return nil
}

// 从文件加载私钥
func loadPrivateKey(filename string) (*rsa.PrivateKey, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	privateKeyBlock, _ := pem.Decode(file)
	if privateKeyBlock == nil {
		return nil, fmt.Errorf("failed to decode PEM block containing the key")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

// 使用公钥加密
func encryptWithPublicKey(plaintext []byte, publicKey *rsa.PublicKey) ([]byte, error) {
	encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plaintext)
	if err != nil {
		return nil, err
	}
	return encryptedBytes, nil
}

// 使用私钥解密
func decryptWithPrivateKey(encryptedBytes []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptedBytes)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

func main() {
	// 生成密钥对
	privateKey, err := generateRSAKeyPair(2048)
	if err != nil {
		fmt.Println("Error generating RSA key pair:", err)
		return
	}

	// 保存私钥
	err = savePrivateKey(privateKey, "privateKey.pem")
	if err != nil {
		fmt.Println("Error saving private key:", err)
		return
	}

	// 加载私钥
	loadedPrivateKey, err := loadPrivateKey("privateKey.pem")
	if err != nil {
		fmt.Println("Error loading private key:", err)
		return
	}

	// 要加密的数据
	data := "Hello, RSA!"

	// 使用公钥加密数据
	encryptedData, err := encryptWithPublicKey([]byte(data), &privateKey.PublicKey)
	if err != nil {
		fmt.Println("Error encrypting data:", err)
		return
	}
	fmt.Println("Encrypted data:", encryptedData)

	// 使用私钥解密数据
	decryptedData, err := decryptWithPrivateKey(encryptedData, loadedPrivateKey)
	if err != nil {
		fmt.Println("Error decrypting data:", err)
		return
	}
	fmt.Println("Decrypted data:", string(decryptedData))
}
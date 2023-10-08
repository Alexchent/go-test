package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
)

func Encrypt(plainText []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plainText))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plainText)

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(ciphertext string, key []byte) ([]byte, error) {
	ciphertextBytes, err := base64.URLEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertextBytes) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	iv := ciphertextBytes[:aes.BlockSize]
	ciphertextBytes = ciphertextBytes[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertextBytes, ciphertextBytes)

	return ciphertextBytes, nil
}

func main() {
	key := []byte("7f60a6645e05382fcef5a4209f0d3d24") // 16 bytes key
	plaintext := []byte("hello world")

	encrypted, err := encrypt(plaintext, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Encrypted:", encrypted)

	decrypted, err := decrypt(encrypted, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Decrypted:", string(decrypted))
}

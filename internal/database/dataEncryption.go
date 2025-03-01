package database

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"os"

	"github.com/joho/godotenv"
)

// load key tu file .env
func loadKey() []byte {
	_ = godotenv.Load() // Load từ .env
	key := os.Getenv("ENCRYPTION_KEY")
	if len(key) != 32 {
		panic("Encryption key must be 32 bytes long")
	}
	return []byte(key)
}

var key = loadKey()

// encrypt ma hoa du lieu
func Encrypt(data string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(data))
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// encrypt giai ma du lieu
func Decrypt(data string) (string, error) {
	ciphertext, _ := base64.URLEncoding.DecodeString(data)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	if len(ciphertext) < aes.BlockSize {
		return "", err
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext), nil
}

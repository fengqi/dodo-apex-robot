package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
)

// AesEncrypt aes-256-cbc pkcs7padding
func AesEncrypt(secretKey, payload string) (string, error) {
	secretKeyBytes, err := hex.DecodeString(secretKey)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(secretKeyBytes)
	if err != nil {
		return "", err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)

	// PKCS7Padding填充
	padding := aes.BlockSize - (len(payload) % aes.BlockSize)
	padText := append([]byte(payload), byte(padding))
	paddedPayload := make([]byte, len(padText))
	copy(paddedPayload, padText)

	encrypted := make([]byte, len(paddedPayload))
	mode.CryptBlocks(encrypted, paddedPayload)

	return hex.EncodeToString(encrypted), nil
}

func AesDecrypt(secretKey, payload string) ([]byte, error) {
	payloadBytes, err := hex.DecodeString(payload)
	if err != nil {
		return nil, err
	}

	secretKeyBytes, err := hex.DecodeString(secretKey)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, 16) // 使用全零IV向量

	block, err := aes.NewCipher(secretKeyBytes)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	result := make([]byte, len(payloadBytes))
	mode.CryptBlocks(result, payloadBytes)

	// PKCS7Padding 填充处理（如果需要）
	unpaddedResult, err := pkcs7Unpad(result)
	if err != nil {
		return nil, err
	}

	return unpaddedResult, nil
}

// pkcs7Pad 对数据进行 PKCS7 填充
func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// pkcs7Unpad 移除 PKCS7 填充
func pkcs7Unpad(data []byte) ([]byte, error) {
	length := len(data)
	unpadding := int(data[length-1])
	if unpadding >= 16 {
		return nil, errors.New("Invalid PKCS7 padding")
	}
	return data[:(length - unpadding)], nil
}

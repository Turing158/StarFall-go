package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// AES的加密思路和代码来自https://blog.csdn.net/qq_35807303/article/details/141527963

var AESKey = []byte("StarFallSecureKeyAndTheKeyMust32")

func AesEncrypt(text string) (string, error) {
	cipherBlock, err := aes.NewCipher(AESKey)
	if err != nil {
		return "", err
	}
	plaintext := padding([]byte(text))

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	blockMode := cipher.NewCBCEncrypter(cipherBlock, iv)
	blockMode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func AesDecrypt(code string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		return "", err
	}

	cipherBlock, err := aes.NewCipher(AESKey)
	if err != nil {
		return "", err
	}

	if len(data) < aes.BlockSize {
		return "", fmt.Errorf("code too short")
	}

	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]

	blockMode := cipher.NewCBCDecrypter(cipherBlock, iv)
	blockMode.CryptBlocks(data, data)

	data = unPadding(data)
	return string(data), nil
}

// pad 对数据进行填充，使其长度为块大小的倍数。
func padding(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

// unpad 移除填充数据。
func unPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

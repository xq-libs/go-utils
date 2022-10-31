package crypt

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

/*
 * PKCSPadding
 */
func PKCS5Padding(data []byte, blockSize int) []byte {
	return PKCSPadding(data, 8)
}

func PKCSPadding(data []byte, blockSize int) []byte {
	paddingLength := blockSize - len(data)%blockSize
	paddingData := bytes.Repeat([]byte{byte(paddingLength)}, paddingLength)
	return append(data, paddingData...)
}

func PKCSUnPadding(data []byte) []byte {
	length := len(data)
	paddingLength := int(data[length-1])
	return data[:(length - paddingLength)]
}

/*
 * DES
 */
func DesEncrypt(data []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	data = PKCSPadding(data, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, iv)
	encryptedData := make([]byte, len(data))
	blockMode.CryptBlocks(encryptedData, data)
	return encryptedData, nil
}

func DesDecrypt(data []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	decryptedData := make([]byte, len(data))
	blockMode.CryptBlocks(decryptedData, data)
	decryptedData = PKCSUnPadding(decryptedData)
	return decryptedData, nil
}

/*
 * AES
 */

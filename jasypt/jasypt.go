package jasypt

import (
	"github.com/xq-libs/go-utils/crypt"

	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"
)

const (
	PBEWithMD5AndDES            = "PBEWithMD5AndDES"
	PBEWITHHMACSHA512ANDAES_256 = "PBEWITHHMACSHA512ANDAES_256"
)

var (
	DefaultSalt       = []byte{}
	DefaultIterations = 1000
)

func init() {
	salt, err := generateSalt(8)
	if err != nil {
		panic(fmt.Sprintf("Fatal error for generate salt: %s \n", err))
	}
	DefaultSalt = salt
}

type Options struct {
	Salt       []byte
	Iterations int
	Password   string
	Algorithm  string
}

type Jasypt interface {
	Encrypt(plainText string) (string, error)
	Decrypt(encryptedText string) (string, error)
}

type FirstJasypt struct {
	Opts Options
}

type SecondJasypt struct {
	Opts Options
}

func GetJasypt(opts Options) Jasypt {
	if opts.Iterations < 1 {
		opts.Iterations = DefaultIterations
	}
	if opts.Algorithm == PBEWithMD5AndDES {
		return &FirstJasypt{Opts: opts}
	} else if opts.Algorithm == PBEWITHHMACSHA512ANDAES_256 {
		return &SecondJasypt{Opts: opts}
	} else {
		return &FirstJasypt{Opts: opts}
	}
}

/*
 * PBEWithMD5AndDES
 */
func (j FirstJasypt) Encrypt(plainText string) (string, error) {
	//1. generate salt
	salt, err := generateSalt(8)
	if err != nil {
		return "", fmt.Errorf("Jasypt error for generate salt: %s \n", err)
	}
	//2. appen padding str
	plainText = appendPadding(plainText, 8)
	//3. generate key
	dk, iv := getDerivedKey(j.Opts.Password, salt, j.Opts.Iterations)
	//4. des encrypt data
	encryptedText, err := crypt.DesEncrypt([]byte(plainText), dk, iv)
	if err != nil {
		return "", nil
	}
	//5. base64 encode
	data := append(salt, encryptedText...)
	return base64.StdEncoding.EncodeToString(data), nil
}

/*
 * PBEWithMD5AndDES
 */
func (j FirstJasypt) Decrypt(encryptedText string) (string, error) {
	//1. base64 decode
	data, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}
	//2. get salt
	salt := data[:8]
	//3. get data
	encryptedData := data[8:]
	//3. generate key
	dk, iv := getDerivedKey(j.Opts.Password, salt, j.Opts.Iterations)
	//4. des decrypt data
	decryptedText, err := crypt.DesDecrypt([]byte(encryptedData), dk, iv)
	if err != nil {
		return "", err
	}
	//5. remove padding str
	return removePadding(string(decryptedText), 8), nil
}

/*
 * PBEWITHHMACSHA512ANDAES_256
 */
func (j SecondJasypt) Encrypt(plainText string) (string, error) {
	return "", nil
}

/*
 * PBEWITHHMACSHA512ANDAES_256
 */
func (j SecondJasypt) Decrypt(encryptedText string) (string, error) {
	return "", nil
}

func getDerivedKey(password string, salt []byte, count int) ([]byte, []byte) {
	key := md5.Sum([]byte(password + string(salt)))
	for i := 0; i < count-1; i++ {
		key = md5.Sum(key[:])
	}
	return key[:8], key[8:]
}

func generateSalt(length int) ([]byte, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	return salt, err
}

func appendPadding(plainText string, blockSize int) string {
	paddingLength := blockSize - len(plainText)%blockSize
	return plainText + strings.Repeat(string(rune(paddingLength)), paddingLength)
}

func removePadding(plainText string, blockSize int) string {
	length := len(plainText)
	paddingLength := int(plainText[length-1])
	return plainText[:(length - paddingLength)]
}

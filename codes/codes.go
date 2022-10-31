package codes

import (
	"encoding/base64"
	"errors"
	"strconv"
	"strings"
)

const (
	CODE32 = "ABCDEFGHIJKLMNOPQRSTUVWXYZ012345"
	CODE58 = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	CODE62 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	CODE64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
)

var (
	code32Map = map[byte]int{}
	code58Map = map[byte]int{}
	code62Map = map[byte]int{}
	code64Map = map[byte]int{}
)

func init() {
	for i := range []byte(CODE32) {
		code32Map[CODE32[i]] = i
	}

	for i := range []byte(CODE58) {
		code58Map[CODE58[i]] = i
	}

	for i := range []byte(CODE62) {
		code62Map[CODE62[i]] = i
	}

	for i := range []byte(CODE64) {
		code64Map[CODE64[i]] = i
	}
}

func EncodeIntBase2(id int64) string {
	return strconv.FormatInt(id, 2)
}

func DecodeIntBase2(str string) (int64, error) {
	return strconv.ParseInt(str, 2, 64)
}

func EncodeIntBase10(id int64) string {
	return strconv.FormatInt(id, 10)
}

func DecodeIntBase10(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

func EncodeIntBase32(id int64) string {
	return encodeIntBase(id, 32, CODE32)
}

func DecodeIntBase32(str string) (int64, error) {
	return decodeIntBase(str, 32, code32Map)
}

func EncodeIntBase58(id int64) string {
	return encodeIntBase(id, 58, CODE58)
}

func DecodeIntBase58(str string) (int64, error) {
	return decodeIntBase(str, 58, code58Map)
}

func EncodeIntBase62(id int64) string {
	return encodeIntBase(id, 62, CODE62)
}

func DecodeIntBase62(str string) (int64, error) {
	return decodeIntBase(str, 62, code62Map)
}

func EncodeIntBase64(id int64) string {
	return encodeIntBase(id, 64, CODE64)
}

func DecodeIntBase64(str string) (int64, error) {
	return decodeIntBase(str, 64, code64Map)
}

func EncodeStrBase64(str string) (string, error) {
	return base64.StdEncoding.EncodeToString([]byte(str)), nil
}

func DecodeStrBase64(str string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	return string(data), err
}

// private methods
func encodeIntBase(id int64, size int64, m string) string {
	if id == 0 {
		return string(m[0])
	}
	r := make([]byte, 0)
	for id > 0 {
		r = append(r, m[id % size])
		id /= size
	}
	return string(reverseArray(r))
}

func decodeIntBase(str string, size int64, m map[byte]int) (int64, error) {
	str = strings.TrimSpace(str)
	if len(str) < 1 {
		return -1, errors.New("invalid string")
	}
	var r int64 = 0
	for i := range []byte(str) {
		r = r * size + int64(m[str[i]])
	}
	return r, nil
}

func reverseArray(arr []byte) []byte {
	for i, j := 0, len(arr) - 1; i < j; i, j = i + 1, j -1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

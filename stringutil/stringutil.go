package stringutil

import (
	"log"
	"strconv"
	"strings"
)

func IsEmpty(str string) bool {
	return len(str) == 0
}

func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}

func IsBlank(str string) bool {
	return IsEmpty(str) || len(strings.TrimSpace(str)) == 0
}

func IsNotBlank(str string) bool {
	return !IsBlank(str)
}

func ToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Panicf("String: %s convert to int failure.", str)
	}
	return i
}

func ToInt64(str string) int64 {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Panicf("String: %s convert to int64 failure.", str)
	}
	return i
}

func ToUInt64(str string) uint64 {
	i, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		log.Panicf("String: %s convert to uint64 failure.", str)
	}
	return i
}

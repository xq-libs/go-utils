package stringutil

import (
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
package mathutil

import "github.com/xq-libs/go-utils/types"

func MaxNumber[T types.Number](a, b T) T {
	if a >= b {
		return a
	}
	return b
}

func MinNumber[T types.Number](a, b T) T {
	if a <= b {
		return a
	}
	return b
}

package types

type SignedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | uint | uint8 | uint16 | uint32 | uint64
}

type UnsignedInteger interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Integer interface {
	SignedInteger | UnsignedInteger
}

type Float interface {
	~float32 | ~float64
}

type Number interface {
	Integer | Float
}

type Function[T, R any] func(T) R

type Function2[T1, T2, R any] func(T1, T2) R

type Supplier[R any] func() R

type Consumer[T any] func(T)

type Consumer2[T1, T2 any] func(T1, T2)

type Predicate[T any] func(T) bool

type Predicate2[T1, T2 any] func(T1, T2) bool

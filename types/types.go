package types

type Function[T, R any] func(T) R

type Function2[T1, T2, R any] func(T1, T2) R

type Supplier[R any] func() R

type Consumer[T any] func(T)

type Consumer2[T1, T2 any] func(T1, T2)

type Predicate[T any] func(T) bool

type Predicate2[T1, T2 any] func(T1, T2) bool

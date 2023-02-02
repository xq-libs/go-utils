package optional

import "github.com/xq-libs/go-utils/types"

type Optional[T any] struct {
	Data    T
	HasData bool
}

func OfEmpty[T any]() Optional[T] {
	return Optional[T]{
		HasData: false,
	}
}

func Of[T any](t T) Optional[T] {
	return Optional[T]{
		Data:    t,
		HasData: true,
	}
}

func Map[T any, R any](o Optional[T], f types.Function[T, R]) Optional[R] {
	if o.HasData {
		return Optional[R]{
			Data:    f(o.Data),
			HasData: true,
		}
	} else {
		return Optional[R]{
			HasData: false,
		}
	}
}

func FlatMap[T any, R any](o Optional[Optional[T]], f types.Function[T, R]) Optional[R] {
	if o.HasData {
		return Map(o.Data, f)
	} else {
		return Optional[R]{
			HasData: false,
		}
	}
}

func (o *Optional[T]) Filter(p types.Predicate[T]) Optional[T] {
	if o.HasData && p(o.Data) {
		return Optional[T]{
			Data:    o.Data,
			HasData: true,
		}
	} else {
		return Optional[T]{
			HasData: false,
		}
	}
}

func (o *Optional[T]) orElse(other T) T {
	if o.HasData {
		return o.Data
	} else {
		return other
	}
}

func (o *Optional[T]) orElseGet(s types.Supplier[T]) T {
	if o.HasData {
		return o.Data
	} else {
		return s()
	}
}

func (o *Optional[T]) ifPresent(c types.Consumer[T]) {
	if o.HasData {
		c(o.Data)
	}
}

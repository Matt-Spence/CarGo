package cargo

type Option[T any] struct {
	val T
	some bool
}

func (o Option[T]) Unwrap() T {
	if o.some {
		return o.val
	}
	panic("unable to unwrap option")
}

func (o Option[T]) UnwrapOr(t T) T {
	if o.some {
		return o.val
	}
	return t
}

func (o Option[T]) IsSome() bool {
	return o.some
}

func (o Option[T]) IsNone() bool {
	return !o.some
}

func (o Option[T]) Expect(str string) T {
	if o.some {
		return o.val
	}
	panic(str)
}

func Some[T any](t T) Option[T] {
	return Option[T]{val: t, some: true}
}

func None() Option[any] {
	return Option[any]{}
}





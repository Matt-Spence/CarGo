package cargo

type Result[T any] struct {
	Value T
	Err   error
}

func (r Result[T]) Unwrap() T {
	return r.Value
}

func (r Result[T]) UnwrapOr(val T) T {
	if r.Err == nil {
		return r.Value
	}
	return val
}

func (r Result[T]) IsOk() bool {
	return r.Err == nil
}

func (r Result[T]) IsErr() bool {
	return r.Err != nil
}

func (r Result[T]) IsOkAnd(f func(T) bool) bool {
	return (r.Err == nil) && f(r.Value)
}

func (r Result[T]) IsErrAnd(f func(error) bool) bool {
	return (r.Err != nil) && f(r.Err)
}

func (r Result[T]) Then(f func(Result[T]) any) any {
	return f(r)
}

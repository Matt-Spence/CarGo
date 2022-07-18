package cargo

import "errors"

type Result[T any] struct {
	Ok T
	Err   error
}

func (r Result[T]) Unwrap() T {
	return r.Ok
}

func (r Result[T]) UnwrapOr(val T) T {
	if r.Err == nil {
		return r.Ok
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
	return (r.Err == nil) && f(r.Ok)
}

func (r Result[T]) IsErrAnd(f func(error) bool) bool {
	return (r.Err != nil) && f(r.Err)
}

func (r Result[T]) Then(f func(Result[T]) any) any {
	return f(r)
}

// convenience functions
func Ok[T any](ok T) Result[T] {
	return Result[T]{Ok: ok}
}

func Err(err string) error {
	return errors.New(err)
}
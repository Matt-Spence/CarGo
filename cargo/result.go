package cargo

type Result[T any] interface {
	Unwrap() T
	UnwrapOr(T) T
	IsOk() bool
	IsErr() bool
	IsOkAnd(func(T) bool) bool
	IsErrAnd(f func() bool) bool
	Then(func(T) any) any
}

type Ok[T any] struct {
	Val T
}

func (r Ok[T]) Unwrap() T {
	return r.Val
}

func (r Ok[T]) UnwrapOr(val T) T {
	return r.Val
}

func (r Ok[T]) IsOk() bool {
	return true
}

func (r Ok[T]) IsErr() bool {
	return false
}

func (r Ok[T]) IsOkAnd(f func(T) bool) bool {
	return f(r.Val)
}

func (r Ok[T]) IsErrAnd(f func() bool) bool {
	return false
}

func (r Ok[T]) Then(f func(T) any) any {
	return f(r.Val)
}

type Err[T any] struct {
	Val T
}

func (r Err[T]) Unwrap() T {
	return r.Val
}

func (r Err[T]) UnwrapOr(val T) T {
	return val
}

func (r Err[T]) IsOk() bool {
	return false
}

func (r Err[T]) IsErr() bool {
	return true
}

func (r Err[T]) IsOkAnd(f func(T) bool) bool {
	return false
}

func (r Err[T]) IsErrAnd(f func() bool) bool {
	return f()
}

func (r Err[T]) Then(f func(T) any) any {
	return f(r.Val)
}
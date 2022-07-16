package result

type Result[T any] struct {
	value T
	err   error
}

func (r Result[T]) isOk() bool {
	return r.err == nil
}

func (r Result[T]) isErr() bool {
	return r.err != nil
}

func (r Result[T]) isOkAnd(f func(T) bool) bool {
	return (r.err == nil) && f(r.value)
}

func (r Result[T]) isErrAnd(f func(error) bool) bool {
	return (r.err != nil) && f(r.err)
}

func (r Result[T]) then(f func(Result[T]) any) any {
	return f(r)
}

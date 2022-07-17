package cargo

import "errors"

type Match[T comparable, U any] struct {
	Value T
	result U
	match string
}

func (m Match[T, U]) Case(test T, statement func() U) Match[T, U]{
    if(m.Value == test && m.match == "") {
		m.result = statement()
		m.match = "true"
	}

	return m
}

func (m Match[T, U]) Done() Result[U] {
	if(m.match == "") {
		return Result[U]{Err: errors.New("comparable value does not match any test cases")}
	}
	return Result[U]{Value: m.result}
}
package cargo

type Match[T comparable, U any] struct {
	Value T
	result U
	match bool
}

func (m Match[T, U]) Case(test T, statement func() U) Match[T, U]{
    if(m.Value == test && !m.match) {
		m.result = statement()
		m.match = true
	}

	return m
}

func (m Match[T, U]) Done() Result[U] {
	if(!m.match) {
		return Err[U]{}
	}
	return Ok[U]{m.result}
}
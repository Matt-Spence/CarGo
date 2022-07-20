package cargo__test

import (
	"testing"
	. "github.com/Matt-Spence/CarGo/cargo"
)

func TestUnwrap(t *testing.T) {
	var sr Ok[float64] = Ok[float64]{6.0}
	if sr.Unwrap() != 6.0 {
		t.Fail()
	}
}

func TestOkUnwrapOr(t *testing.T) {
	var sr Ok[float64] = Ok[float64]{6.0}
	if sr.UnwrapOr(1.0) != 6.0 {
		t.Fail()
	}
}

func TestErrUnwrapOr(t *testing.T) {
	var sr Err[float64] = Err[float64]{0.0}
	if sr.UnwrapOr(1.0) != 1.0 {
		t.Fail()
	}
}

func TestOkIsOkAnd(t *testing.T) {
	var sr Ok[float64] = Ok[float64]{6.0}
	if !sr.IsOkAnd(func(f float64) bool { return f > 5.0 }) {
		t.Fail()
	}
}

func TestErrIsOkAnd(t *testing.T) {
	var sr Result[float64] = Err[float64]{6.0}
	if sr.IsOkAnd(func(f float64) bool { return f > 5.0 }) {
		t.Fail()
	}
}

func TestOkIsOk(t *testing.T) {
	var sr Result[float64] = Ok[float64]{6.0}
	if !sr.IsOk() {
		t.Fail()
	}
}

func TestErrIsOk(t *testing.T) {
	var sr Result[float64] = Err[float64]{6.0}
	if sr.IsOk() {
		t.Fail()
	}
}

func TestOkIsErrAnd(t *testing.T) {
	var sr Result[float64] = Ok[float64]{6.0}
	if sr.IsErrAnd(func() bool { return true }) {
		t.Fail()
	}
}

func TestErrIsErrAnd(t *testing.T) {
	var sr Result[float64] = Err[float64]{6.0}
	if !sr.IsErrAnd(func() bool { return true }) {
		t.Fail()
	}
}

func TestErrIsErr(t *testing.T) {
	var sr Result[float64] = Err[float64]{6.0}
	if !sr.IsErr() {
		t.Fail()
	}
}

func TestOkThen(t *testing.T) {
	var sr Result[float64] = Ok[float64]{6.0}
	if sr.Then(func(f float64) any {return f + 1.0}) != 7.0 {
		t.Fail()
	}
}

func TestErrThen(t *testing.T) {
	var sr Result[float64] = Err[float64]{6.0}
	if sr.Then(func(f float64) any {return f + 1.0}) != 7.0 {
		t.Fail()
	}
}
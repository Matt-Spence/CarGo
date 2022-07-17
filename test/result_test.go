package cargo__test

import (
	"errors"
	"testing"
	. "github.com/Matt-Spence/CarGo/cargo"
)

func TestIsOkAnd(t *testing.T) {
	var sr Result[float64] = Result[float64]{Value: 6.0, Err: nil}
	if !sr.IsOkAnd(func(f float64) bool { return f > 5.0 }) {
		t.Fail()
	}
}

func TestIsOk(t *testing.T) {
	var sr Result[float64] = Result[float64]{Value: 6.0, Err: nil}
	if !sr.IsOk() {
		t.Fail()
	}
}

func TestIsErrAnd(t *testing.T) {
	var sr Result[float64] = Result[float64]{Value: 6.0, Err: errors.New("error")}
	if !sr.IsErrAnd(func(e error) bool { return e.Error() == "error" }) {
		t.Fail()
	}
}

func TestIsErr(t *testing.T) {
	var sr Result[float64] = Result[float64]{Value: 6.0, Err: errors.New("error")}
	if !sr.IsErr() {
		t.Fail()
	}
}

func TestThen(t *testing.T) {
	var sr Result[float64] = Result[float64]{Value: 6.0, Err: errors.New("error")}
	if sr.Then(func(r Result[float64]) any {return r.Value + 1.0}) != 7.0 {
		t.Fail()
	}
}
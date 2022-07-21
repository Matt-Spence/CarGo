package cargo__test

import (
	"testing"
	. "github.com/Matt-Spence/CarGo/cargo"
)

func TestSomeIsSome(t *testing.T) {
	var op Option[float64] = Some(6.0)
	if !op.IsSome() {
		t.Fail()
	}
}

func TestNoneIsSome(t *testing.T) {
	var op Option[bool] = None[bool]()
	if op.IsSome() {
		t.Fail()
	}
}

func TestSomeIsNone(t *testing.T) {
	var op Option[float64] = Some(6.0)
	if op.IsNone() {
		t.Fail()
	}
}

func TestNoneIsNone(t *testing.T) {
	var op Option[bool] = None[bool]()
	if !op.IsNone() {
		t.Fail()
	}
}

func TestNoneExpect(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()
	None[bool]().Expect("test")
}

func TestSomeExpect(t *testing.T) {
	if Some(1).Expect("test") != 1 {
		t.Fail()
	}
}
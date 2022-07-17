package cargo__test

import (
	"testing"

	. "github.com/Matt-Spence/CarGo/cargo"
)

func TestMatchSingleCase(t *testing.T) {
	var str string = "test"
	var passed Result[bool] = Match[string, bool]{Value: str}.Case("test", func() bool {return true},).Done()

	if passed.IsErr() {
		t.Fail()
	} 
}

func TestMatchMultiCase(t *testing.T) {
	var str string = "test"
	var passed Result[int] = Match[string, int]{Value: str}.
		Case("notThisOne", func() int {return 1}).
		Case("test", func() int {return 2}).
	Done()

	if passed.IsOkAnd(func(i int) bool {return i == 2}) {
		t.Fail()
	} 
}
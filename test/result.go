package test

import (
	"testing"
)

func TestIsOkAnd(t *testing.T) {
	var result Result[float64] = Result[float64]{value: 5.0, err: nil}
	return sr.isOkAnd(func(f float64) bool { return f > 5.0 })
}

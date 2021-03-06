package tst

import (
	"testing"
)

// Blame calls t.Blame with conventional error format.
// Arguments:
// fn - failed function name.
// res - failed function return value.
// expected - expected return value.
func Blame(t *testing.T, fn string, res, expected interface{}) {
	t.Errorf("%s()=>%v (want %v)", fn, res, expected)
}

// CheckError calls Errorf if "res != expected".
func CheckError(t *testing.T, expr string, res, expected interface{}) {
	if res != expected {
		Blame(t, expr, res, expected)
	}
}

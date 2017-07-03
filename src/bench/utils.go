package bench

import (
	"assert"
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
	"time"
)

// Used for benchmark result assignment
// (to disappoint Go optimizer).
var gScore int

// Default random generator for input data.
var random = rand.New(rand.NewSource(time.Now().Unix()))

// Run executes given function b.N times
// and collects the result (score).
// Score pointer is passed into benchmarking function.
func Run(b *testing.B, bench func(*int)) {
	score := 0
	for i := 0; i < b.N; i++ {
		bench(&score)
	}
	gScore = score
}

// RandIntSlice generates slice of 50 random int elements.
func RandIntSlice() []int {
	res, ok := quick.Value(reflect.TypeOf([]int{}), random)
	assert.True(ok)
	return res.Interface().([]int)
}

// RandInt64Slice generates slice of 50 random int64 elements.
func RandInt64Slice() []int64 {
	res, ok := quick.Value(reflect.TypeOf([]int64{}), random)
	assert.True(ok)
	return res.Interface().([]int64)
}

// RandFloat64Slice generates slice of 50 random float64 elements.
func RandFloat64Slice() []float64 {
	res, ok := quick.Value(reflect.TypeOf([]float64{}), random)
	assert.True(ok)
	return res.Interface().([]float64)
}

// RandStringSlice generates slice of 50 random string elements.
func RandStringSlice() []string {
	res, ok := quick.Value(reflect.TypeOf([]string{}), random)
	assert.True(ok)
	return res.Interface().([]string)
}

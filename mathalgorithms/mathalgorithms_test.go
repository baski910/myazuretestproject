package mathalgorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFibonacciIterator(t *testing.T) {
	iter := NewFibonacciIterator()
	assert.Panics(t, func() { iter.Value() }, "Value did not panic before Next")
}

func TestNext(t *testing.T) {
	iter := NewFibonacciIterator()
	iter.Next()
	assert.Equal(t, int32(1), iter.Value())
	iter.Next()
	assert.Equal(t, int32(1), iter.Value())
}

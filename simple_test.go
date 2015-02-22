package main

import (
	"testing"

	"github.com/stvp/assert"
)

// TestEqualsWorks showcases a simple assertion library "assert".
// assert provides a few basic checks and builds upon the regular testing
// framework of Go.
func TestEqualsWorks(t *testing.T) {
	expected := true
	result := true

	assert.Equal(t, expected, result)
}

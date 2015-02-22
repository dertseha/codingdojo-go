package main

import (
	"testing"

	check "gopkg.in/check.v1"
)

// Test initializes the "check" testing framework.
// This framework provides and runs test suites. Per package only one
// main "Test" function is required that starts the framework.
func Test(t *testing.T) { check.TestingT(t) }

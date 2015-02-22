package main

import (
	check "gopkg.in/check.v1"
)

type SimpleSuite struct {
	testObject string
}

// check.Suite registers a new object as a suite. This suite object
// can provide "Test*" methods and further helper such as "SetUpTest".
var _ = check.Suite(&SimpleSuite{})

func (suite *SimpleSuite) SetUpTest(c *check.C) {
	suite.testObject = "1234"
}

func (suite *SimpleSuite) TestChecker(c *check.C) {
	result := suite.testObject

	c.Assert(result, check.Equals, "1234")
}

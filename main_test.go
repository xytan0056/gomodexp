package main

import (
	"testing"
	c "github.com/xytan0056/depwithgen/.gen/client"
)

func TestSomeTest(t *testing.T) {
	t.Logf("testing a test")
	t.Log(c.Client("007"))
}
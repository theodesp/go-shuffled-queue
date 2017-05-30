package shuffled_queue

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) {
	TestingT(t)
}

type MySuite struct{}

var _ = Suite(&MySuite{})

// Smoke test
func (s *MySuite) TestSmoke(c *C) {
	c.Assert(true, Equals, true)
}

// Default Constructor test
func (s *MySuite) TestNewSPQ(c *C) {
	queue := NewSPQ()
	c.Assert(queue.priorities.String(), Equals, "Set{}")
}

func (s *MySuite) BenchmarkNewSPQ(c *C) {
	for i := 0; i < c.N; i++ {
		NewSPQ()
	}
}

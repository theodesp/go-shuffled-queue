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
	c.Assert(len(queue.priorities), Equals, 0)
}

// Add method
func (s *MySuite) TestAdd(c *C) {
	queue := NewSPQ()

	c.Assert(queue.length, Equals, uint(0))

	queue.Add("world")
	queue.Add("world")

	c.Assert(queue.length, Equals, uint(1))
	c.Assert(queue.priorities[0].ToSlice(), DeepEquals, []interface{}{"world"})
}

// AddWithPriority method
func (s *MySuite) TestAddWithPriority(c *C) {
	queue := NewSPQ()

	c.Assert(queue.length, Equals, uint(0))

	queue.AddWithPriority("welt", uint(0))
	queue.AddWithPriority("hello", uint(1))
	queue.AddWithPriority("hello", uint(1))

	c.Assert(queue.length, Equals, uint(2))
}

func (s *MySuite) BenchmarkNewSPQ(c *C) {
	for i := 0; i < c.N; i++ {
		NewSPQ()
	}
}

func (s *MySuite) BenchmarkAdd(c *C) {
	spq := NewSPQ()
	for i := 0; i < c.N; i++ {
		spq.Add(i)
	}
}

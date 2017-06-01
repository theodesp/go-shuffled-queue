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

// Test Default Constructor test.
func (s *MySuite) TestNewSPQ(c *C) {
	queue := NewSPQ()
	c.Assert(queue.length, Equals, uint(0))
}

// Test Add method.
func (s *MySuite) TestAdd(c *C) {
	queue := NewSPQ()

	queue.Add("world")
	queue.Add("world")

	c.Assert(queue.length, Equals, uint(1))
	c.Assert(queue.priorities[0].ToSlice(), DeepEquals, []interface{}{"world"})
}

// Test AddWithPriority method.
func (s *MySuite) TestAddWithPriority(c *C) {
	queue := NewSPQ()

	queue.AddWithPriority("welt", 0)
	queue.AddWithPriority("hello", 1)
	queue.AddWithPriority("hello", 1)

	c.Assert(queue.length, Equals, uint(2))
}

// Test Remove When spq is empty method.
func (s *MySuite) TestRemoveWhenEmpty(c *C) {
	spq := NewSPQ()

	c.Assert(spq.Remove("hello"), Equals, false)
}

// Test Remove Item that do not exist.
func (s *MySuite) TestRemoveWhenItemNotExists(c *C) {
	spq := NewSPQ()

	spq.AddWithPriority("welt", 0)

	c.Assert(spq.Remove("hello"), Equals, false)
}

// Test Remove Item that exists and its the only one.
func (s *MySuite) TestRemoveWhenItemExistsAndOnlyOne(c *C) {
	spq := NewSPQ()

	spq.AddWithPriority("welt", 0)
	spq.AddWithPriority("hello", 0)

	c.Assert(spq.Remove("welt"), Equals, true)

	priority, found := spq.Find("welt")

	c.Assert(priority, Equals, -1)
	c.Assert(found, Equals, false)
}

// Test Remove Item that exists and its not the only one.
func (s *MySuite) TestRemoveWhenItemExistsAndNotOnlyOne(c *C) {
	spq := NewSPQ()

	spq.AddWithPriority("welt", 0)
	spq.AddWithPriority("hello", 1)
	spq.AddWithPriority("hello", 0)

	c.Assert(spq.Remove("hello"), Equals, true)

	// It should exist another one with the same value
	priority, found := spq.Find("hello")

	c.Assert(priority, Equals, 1)
	c.Assert(found, Equals, true)
}

// Test if Removing an Item that is the last in a priority bucket compresses the priority map size.
func (s *MySuite) TestRemoveWhenRemovesEmptyPriorityKeyBuckets(c *C) {
	spq := NewSPQ()

	spq.AddWithPriority("welt", 0)
	spq.Remove("welt")

	c.Assert(spq.length, Equals, uint(0))
}

// Test First returns the first element if its the only one. Does not mutate the queue.
func (s *MySuite) TestFirst(c *C) {
	spq := NewSPQ()

	spq.AddWithPriority("welt", -1)
	spq.AddWithPriority("world", -1)
	spq.AddWithPriority("mold", -1)
	spq.AddWithPriority("hello", -1)
	spq.AddWithPriority("Atme", -1)
	spq.First()
}

// Benchmarks
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

func (s *MySuite) BenchmarkAddWithPriority(c *C) {
	spq := NewSPQ()

	for i := 0; i < c.N; i++ {
		spq.AddWithPriority(i, i)
	}
}

func (s *MySuite) BenchmarkRemoveWhenEmpty(c *C) {
	spq := NewSPQ()

	for i := 0; i < c.N; i++ {
		spq.Remove("hello")
	}
}

func (s *MySuite) BenchmarkRemoveWhenItemNotExists(c *C) {
	spq := NewSPQ()
	spq.AddWithPriority("welt", 0)

	for i := 0; i < c.N; i++ {
		spq.Remove("hello")
	}
}

func (s *MySuite) BenchmarkRemoveWhenItemExistsAndOnlyOne(c *C) {
	spq := NewSPQ()
	spq.AddWithPriority("welt", 0)
	spq.AddWithPriority("hello", 0)
	spq.AddWithPriority("hello", 1)

	for i := 0; i < c.N; i++ {
		spq.Remove("hello")
	}
}

package go_shuffled_queue

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

func contains(Array []string, Item string) bool {
	for _, I := range Array {
		if I == Item {
			return true
		}
	}
	return false
}

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
func (s *MySuite) TestAddPriority(c *C) {
	queue := NewSPQ()

	queue.AddPriority("welt", 0)
	queue.AddPriority("hello", 1)
	queue.AddPriority("hello", 1)

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

	spq.AddPriority("welt", 0)

	c.Assert(spq.Remove("hello"), Equals, false)
}

// Test Remove Item that exists and its the only one.
func (s *MySuite) TestRemoveWhenItemExistsAndOnlyOne(c *C) {
	spq := NewSPQ()

	spq.AddPriority("welt", 0)
	spq.AddPriority("hello", 0)

	c.Assert(spq.Remove("welt"), Equals, true)

	priority, found := spq.FindPriority("welt")

	c.Assert(priority, Equals, -1)
	c.Assert(found, Equals, false)
}

// Test Remove Item that exists and its not the only one.
func (s *MySuite) TestRemoveWhenItemExistsAndNotOnlyOne(c *C) {
	spq := NewSPQ()

	spq.AddPriority("welt", 0)
	spq.AddPriority("hello", 1)
	spq.AddPriority("hello", 0)

	c.Assert(spq.Remove("hello"), Equals, true)

	// It should exist another one with the same value
	priority, found := spq.FindPriority("hello")

	c.Assert(priority, Equals, 1)
	c.Assert(found, Equals, true)
}

// Test if Removing an Item that is the last in a priority bucket compresses the priority map size.
func (s *MySuite) TestRemoveWhenRemovesEmptyPriorityKeyBuckets(c *C) {
	spq := NewSPQ()

	spq.AddPriority("welt", 0)
	spq.Remove("welt")

	c.Assert(spq.length, Equals, uint(0))
}

// Test First on empty queue
func (s *MySuite) TestFirstOnEmptyQueue(c *C) {
	spq := NewSPQ()

	item, ok := spq.First()

	c.Assert(item, IsNil)
	c.Assert(ok, Equals, false)
}

// Test First does not mutate the queue
func (s *MySuite) TestFirstDoesNotMutate(c *C) {
	spq := NewSPQ()

	spq.AddPriority("welt", -1)
	spq.AddPriority("world", -2)
	spq.AddPriority("mold", -2)
	spq.AddPriority("hello", -2)
	spq.AddPriority("Atme", -3)

	c.Assert(spq.length, Equals, uint(5))

	_, ok := spq.First()

	c.Assert(ok, Equals, true)
	c.Assert(spq.length, Equals, uint(5))
}


// Test First returns the highest priority item if its the only one with the same priority.
func (s *MySuite) TestFirstOnlyOne(c *C) {
	spq := NewSPQ()

	spq.AddPriority("welt", -1)
	spq.AddPriority("world", -2)
	spq.AddPriority("mold", -2)
	spq.AddPriority("hello", -2)
	spq.AddPriority("Atme", -3)

	item, ok := spq.First()

	c.Assert(ok, Equals, true)
	c.Assert(contains([]string{"Atme"}, item.(string)), Equals, true)
	c.Assert(contains([]string{"world", "mold", "hello"}, item.(string)), Equals, false)
	c.Assert(contains([]string{"welt"}, item.(string)), Equals, false)
}

// Test First returns the a random highest priority item from the bucket of items with the same priority.
func (s *MySuite) TestFirstPicksRandomWithSamePriority(c *C) {
	spq := NewSPQ()

	spq.AddPriority("welt", -1)
	spq.AddPriority("world", -2)
	spq.AddPriority("mold", -2)
	spq.AddPriority("hello", -2)
	spq.AddPriority("Atme", -1)

	item, ok := spq.First()

	c.Assert(ok, Equals, true)
	c.Assert(contains([]string{"world", "mold", "hello"}, item.(string)), Equals, true)
	c.Assert(contains([]string{"Atme"}, item.(string)), Equals, false)
	c.Assert(contains([]string{"welt"}, item.(string)), Equals, false)
}

// Test Last on empty queue
func (s *MySuite) TestLastOnEmptyQueue(c *C) {
	spq := NewSPQ()

	item, ok := spq.Last()

	c.Assert(item, IsNil)
	c.Assert(ok, Equals, false)
}

// Test Last does not mutate the queue
func (s *MySuite) TestLastDoesNotMutate(c *C) {
	spq := NewSPQ()

	spq.AddPriority("welt", -1)
	spq.AddPriority("world", -2)
	spq.AddPriority("mold", -2)
	spq.AddPriority("hello", -2)
	spq.AddPriority("Atme", -3)

	c.Assert(spq.length, Equals, uint(5))

	_, ok := spq.Last()

	c.Assert(ok, Equals, true)
	c.Assert(spq.length, Equals, uint(5))
}


// Test Last returns the lowest priority item if its the only one with the same priority.
func (s *MySuite) TestLastOnlyOne(c *C) {
	spq := NewSPQ()

	spq.AddPriority("welt", -1)
	spq.AddPriority("world", -2)
	spq.AddPriority("mold", -2)
	spq.AddPriority("hello", -2)
	spq.AddPriority("Atme", -3)

	item, ok := spq.Last()

	c.Assert(ok, Equals, true)
	c.Assert(contains([]string{"welt"}, item.(string)), Equals, true)
	c.Assert(contains([]string{"Atme"}, item.(string)), Equals, false)
	c.Assert(contains([]string{"world", "mold", "hello"}, item.(string)), Equals, false)

}

// Test Last returns the a random lowest priority item from the bucket of items with the same priority.
func (s *MySuite) TestLastPicksRandomWithSamePriority(c *C) {
	spq := NewSPQ()

	spq.AddPriority("welt", -1)
	spq.AddPriority("world", -2)
	spq.AddPriority("mold", -2)
	spq.AddPriority("hello", -2)
	spq.AddPriority("Atme", -1)

	item, ok := spq.Last()

	c.Assert(ok, Equals, true)
	c.Assert(contains([]string{"welt", "Atme"}, item.(string)), Equals, true)
	c.Assert(contains([]string{"world", "mold", "hello"}, item.(string)), Equals, false)
}


// Test Pop on empty queue
func (s *MySuite) TestPopOnEmptyQueue(c *C) {
	spq := NewSPQ()

	item, ok := spq.Pop()

	c.Assert(item, IsNil)
	c.Assert(ok, Equals, false)
}

// Test Pop mutates the queue
func (s *MySuite) TestPopDoesMutate(c *C) {
	spq := NewSPQ()

	spq.AddPriority("welt", -1)
	spq.AddPriority("world", -2)
	spq.AddPriority("mold", -2)
	spq.AddPriority("hello", -2)
	spq.AddPriority("Atme", -3)

	c.Assert(spq.length, Equals, uint(5))

	_, ok := spq.Pop()

	c.Assert(ok, Equals, true)
	c.Assert(spq.length, Equals, uint(4))
}


// Test Pop returns the lowest priority item if its the only one with the same priority.
func (s *MySuite) TestPopOnlyOne(c *C) {
	spq := NewSPQ()

	spq.AddPriority("welt", -1)
	spq.AddPriority("world", -2)
	spq.AddPriority("mold", -2)
	spq.AddPriority("hello", -2)
	spq.AddPriority("Atme", -3)

	item, ok := spq.Pop()

	c.Assert(ok, Equals, true)
	c.Assert(contains([]string{"welt"}, item.(string)), Equals, true)
	c.Assert(contains([]string{"Atme"}, item.(string)), Equals, false)
	c.Assert(contains([]string{"world", "mold", "hello"}, item.(string)), Equals, false)

}

// Test Pop returns the a random lowest priority item from the bucket of items with the same priority.
func (s *MySuite) TestPopPicksRandomWithSamePriority(c *C) {
	spq := NewSPQ()

	spq.AddPriority("welt", -1)
	spq.AddPriority("world", -2)
	spq.AddPriority("mold", -2)
	spq.AddPriority("hello", -2)
	spq.AddPriority("Atme", -1)

	item, ok := spq.Pop()

	c.Assert(ok, Equals, true)
	c.Assert(contains([]string{"welt", "Atme"}, item.(string)), Equals, true)
	c.Assert(contains([]string{"world", "mold", "hello"}, item.(string)), Equals, false)
}

// Test Shift on empty queue
func (s *MySuite) TestShiftOnEmptyQueue(c *C) {
	spq := NewSPQ()

	item, ok := spq.Shift()

	c.Assert(item, IsNil)
	c.Assert(ok, Equals, false)
}

// Test Shift mutates the queue
func (s *MySuite) TestShiftDoesMutate(c *C) {
	spq := NewSPQ()

	spq.AddPriority("welt", -1)
	spq.AddPriority("world", -2)
	spq.AddPriority("mold", -2)
	spq.AddPriority("hello", -2)
	spq.AddPriority("Atme", -3)

	c.Assert(spq.length, Equals, uint(5))

	_, ok := spq.Shift()

	c.Assert(ok, Equals, true)
	c.Assert(spq.length, Equals, uint(4))
}


// Test Shift returns the highest priority item if its the only one with the same priority.
func (s *MySuite) TestShiftOnlyOne(c *C) {
	spq := NewSPQ()

	spq.AddPriority("welt", -1)
	spq.AddPriority("world", -2)
	spq.AddPriority("mold", -2)
	spq.AddPriority("hello", -2)
	spq.AddPriority("Atme", -3)

	item, ok := spq.Shift()

	c.Assert(ok, Equals, true)
	c.Assert(contains([]string{"Atme"}, item.(string)), Equals, true)
	c.Assert(contains([]string{"world", "mold", "hello"}, item.(string)), Equals, false)
	c.Assert(contains([]string{"welt"}, item.(string)), Equals, false)
}

// Test Shift returns the a random highest priority item from the bucket of items with the same priority.
func (s *MySuite) TestShiftPicksRandomWithSamePriority(c *C) {
	spq := NewSPQ()

	spq.AddPriority("welt", -1)
	spq.AddPriority("world", -2)
	spq.AddPriority("mold", -2)
	spq.AddPriority("hello", -2)
	spq.AddPriority("Atme", -1)

	item, ok := spq.Shift()

	c.Assert(ok, Equals, true)
	c.Assert(contains([]string{"world", "mold", "hello"}, item.(string)), Equals, true)
	c.Assert(contains([]string{"Atme"}, item.(string)), Equals, false)
	c.Assert(contains([]string{"welt"}, item.(string)), Equals, false)
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

func (s *MySuite) BenchmarkAddPriority(c *C) {
	spq := NewSPQ()

	for i := 0; i < c.N; i++ {
		spq.AddPriority(i, i)
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
	spq.AddPriority("welt", 0)

	for i := 0; i < c.N; i++ {
		spq.Remove("hello")
	}
}

func (s *MySuite) BenchmarkRemoveWhenItemExistsAndOnlyOne(c *C) {
	spq := NewSPQ()
	spq.AddPriority("welt", 0)
	spq.AddPriority("hello", 0)
	spq.AddPriority("hello", 1)

	for i := 0; i < c.N; i++ {
		spq.Remove("hello")
	}
}

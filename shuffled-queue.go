/*
Open Source Initiative OSI - The MIT License (MIT):Licensing

The MIT License (MIT)
Copyright (c) 2017 Theo Despoudis (thdespou@hotmail.com)

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Package shuffled_queue implements a non thread safe priority queue that shuffles elements with the same priority.

package shuffled_queue

import (
	"github.com/deckarep/golang-set"
	"sort"
	"math/rand"
	"time"
)

// The default priority of all items unless specified otherwise
const DefaultPriority int = 0

type ShuffledPriorityQueue struct {
	priorities map[int]mapset.Set
	keys       []int
	length     uint
}

// Creates and returns a reference to an empty shuffled priority queue.
func NewSPQ() *ShuffledPriorityQueue {
	spq := &ShuffledPriorityQueue{
		priorities: make(map[int]mapset.Set),
		keys: []int{},
		length:     uint(0)}

	return spq
}

// Adds an item to the priority queue using the default priority.
// Returns the value added.
func (spq *ShuffledPriorityQueue) Add(Value interface{}) interface{} {
	return spq.AddWithPriority(Value, DefaultPriority)
}

// Adds an item to the priority queue using a specified priority.
// Returns the value added.
func (spq *ShuffledPriorityQueue) AddWithPriority(Value interface{}, Priority int) interface{} {
	_, ok := spq.priorities[Priority]

	if !ok {
		spq.priorities[Priority] = mapset.NewSet()
		spq.keys = append(spq.keys, Priority)

		// We maintain a sorted list of keys for Pop, Shift operations
		sort.Ints(spq.keys)
	}

	if spq.priorities[Priority].Add(Value) {
		spq.length += 1
	}

	return Value
}

// Remove the item from the queue if exists.
// Returns true if item was removed or false if the item was not found.
func (spq *ShuffledPriorityQueue) Remove(Value interface{}) bool {
	priority, found := spq.FindPriority(Value)

	if !found {
		return false
	}

	spq.priorities[priority].Remove(Value)

	// Cleanup the priority queue so that it does not grow too big
	if spq.priorities[priority].Cardinality() == 0 {
		spq.removePriorityKey(priority)
	}

	return true
}

// Attempts to find the first specified item and returns its priority.
// Returns true if found otherwise false.
func (spq *ShuffledPriorityQueue) FindPriority(Value interface{}) (int, bool) {
	if spq.length == 0 {
		return -1, false
	}

	for i := 0; i < len(spq.keys); i += 1 {
		priority := spq.keys[i];

		if spq.priorities[priority].Contains(Value) {
			// First found first served
			return priority, true
		}
	}

	return -1, false
}

// Returns the first item from the queue if its the only one.
// Returns true if found otherwise false.
func (spq *ShuffledPriorityQueue) First() (interface{}, bool) {
	if spq.length == 0 {
		return nil, false
	}

	// We assume keys are sorted otherwise we sort them now
	if !sort.IntsAreSorted(spq.keys) {
		sort.Ints(spq.keys)
	}

	lowestPriorityKey := spq.keys[0]

	item := spq.pickRandom(spq.priorities[lowestPriorityKey])
	return item, true
}

// Returns the last item from the queue if its the only one.
// Returns true if found otherwise false. Does not mutate the queue.
func (spq *ShuffledPriorityQueue) Last() (interface{}, bool) {
	if spq.length == 0 {
		return nil, false
	}

	// We assume keys are sorted otherwise we sort them now
	if !sort.IntsAreSorted(spq.keys) {
		sort.Ints(spq.keys)
	}

	highestPriorityKey := spq.keys[len(spq.keys) - 1]

	item := spq.pickRandom(spq.priorities[highestPriorityKey])
	return item, true
}

// Removes and returns the highest priority item from the queue if its the only one.
// Returns true if found otherwise false.
func (spq *ShuffledPriorityQueue) Pop() (interface{}, bool) {
	if spq.length == 0 {
		return nil, false
	}

	item, _ := spq.Last()
	return item, spq.Remove(item)
}

// Removes and returns the lowest priority item from the queue if its the only one.
// Returns true if found otherwise false.
func (spq *ShuffledPriorityQueue) Shift() (interface{}, bool) {
	if spq.length == 0 {
		return nil, false
	}

	item, _ := spq.First()
	return item, spq.Remove(item)
}


// Picks a random element from the set
func (spq *ShuffledPriorityQueue) pickRandom(S mapset.Set) interface{} {
	rand.Seed(time.Now().UTC().UnixNano())
	randomIndex := rand.Intn(S.Cardinality())

	return S.ToSlice()[randomIndex]
}

func (spq *ShuffledPriorityQueue) removePriorityKey(Priority int) {
	delete(spq.priorities, Priority)
	sort.Ints(spq.keys)
	i := sort.SearchInts(spq.keys, Priority)

	// https://github.com/golang/go/wiki/SliceTricks#Delete
	spq.keys = append(spq.keys[:i], spq.keys[i + 1:]...)
	spq.length -= 1
}


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

// Package shuffled_queue implements a priority queue that shuffles elements with the same priority.

package shuffled_queue

import (
	"github.com/deckarep/golang-set"
)

// The default priority of all items unless specified otherwise
const DefaultPriority uint = uint(0)

type ShuffledPriorityQueue struct {
	priorities map[uint]mapset.Set
	length     uint
}

// Creates and returns a reference to an empty shuffled priority queue.
func NewSPQ() *ShuffledPriorityQueue {
	spq := &ShuffledPriorityQueue{
		priorities: make(map[uint]mapset.Set),
		length:     0}

	return spq
}

// Adds an item to the priority queue using the default priority.
// Returns the value added.
func (spq *ShuffledPriorityQueue) Add(Value interface{}) interface{} {
	return spq.AddWithPriority(Value, DefaultPriority)
}

// Adds an item to the priority queue using a specified priority.
// Returns the value added.
func (spq *ShuffledPriorityQueue) AddWithPriority(Value interface{}, priority uint) interface{} {
	_, ok := spq.priorities[priority]

	if !ok {
		spq.priorities[priority] = mapset.NewSet()
	}

	if spq.priorities[priority].Add(Value) {
		spq.length += 1
	}

	return Value
}

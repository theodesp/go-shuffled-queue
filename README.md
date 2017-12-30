# go-shuffled-queue
A priority queue that shuffles elements with the same priority written in Go 
and inspired by [shuffled-priority-queue](https://github.com/mafintosh/shuffled-priority-queue).

## Install

`$ go get -u github.com/theodesp/go-shuffled-queue`

## Usage
```go
queue := shuffledQueue.NewSPQ()

queue.Add("hello") // Default Priority is 0
queue.Add("world") // Default Priority is 0

queue.AddWithPriority("welt", 1)
queue.AddWithPriority("verden", 2)
queue.AddWithPriority("verden", 3)


fmt.Println(queue.Pop()) // returns "verden", true
fmt.Println(queue.Pop()) // returns "verden", true
fmt.Println(queue.Pop()) // returns "welt", true
fmt.Println(queue.Pop()) // returns "hello", true or "world", true
fmt.Println(queue.Pop()) // returns "hello", true or "world", true
fmt.Println(queue.Pop()) // returns nil, false

```


## API

#### `queue := shuffledQueue.NewSPQ()`
Create a new queue.


#### `value := queue.Add(value)`

Add a new value to the queue. Accepts single values. The value is returned for convenience. It also assigns it with a default priority.

#### `value := queue.AddWithPriority(value)`

Add a new PriorityQueueItem to the queue. The value is returned for convenience.


#### `value := queue.Remove(value)`

Remove a value from the queue.


#### `value := queue.Pop()`

Pop the value with the highest priority off the queue. If multiple values have the same priority a random one is popped.
If the queue is empty it will return a value of nil.

#### `value := queue.Last()`

Same as Pop() but does not mutate the queue.

#### `value := queue.Shift()`

Same as Pop() but returns a value with the lowest priority.

#### `value := queue.First()`

Same as Shift() but does not mutate the queue.


## Licence
MIT @ 2017

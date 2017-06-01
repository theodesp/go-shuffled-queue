# go-shuffled-queue
A priority queue that shuffles elements with the same priority written in Go 
and inspired by [shuffled-priority-queue](https://github.com/mafintosh/shuffled-priority-queue).

## WIP
Work in progress

## Install

`$ go get github.com/theodesp/go-shuffled-queue`

## Usage
```go
queue := shuffledQueue.NewSPQ()

queue.add("hello") // Default Priority is 0

queue.add("world") // Default Priority is 0

queue.addWithPriority(&{priority: 1, value: "welt"})

queue.addWithPriority(&{priority: 2, value: "verden"})


fmt.Println(queue.pop()) // returns {value: 'verden'}
fmt.Println(queue.pop()) // returns {value: 'verden'}
fmt.Println(queue.pop()) // returns {value: 'welt'}
fmt.Println(queue.pop()) // returns {value: 'hello'} or {value: 'world'}
fmt.Println(queue.pop()) // returns {value: 'hello'} or {value: 'world'}
fmt.Println(queue.pop()) // returns {value: nil} empty queue

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

Same as pop() but does not mutate the queue.

#### `value := queue.Shift()`

Same as pop() but returns a value with the lowest priority.

#### `value := queue.First()`

Same as shift() but does not mutate the queue.


#### `it := queue.Iterator()`

Returns an Iterator object that you can
use to range over the queue items from lowest priority to highest. For example:

```go
type YourType struct {
    Value string
}

it := queue.Iterator()

for elem := range it.C {
        fmt.Printf("Job %+v\n", elem.(*YourType).Value)
    }
```

#### `it := queue.ReverseIterator()`

Returns an Iterator object that you can
use to range over the queue items from highest priority to the lowest. For example:

```go
type YourType struct {
    Value string
}

it := set.ReverseIterator()

for elem := range it.C {
        fmt.Printf("Job %+v\n", elem.(*YourType).Value)
    }
```

## Licence
MIT

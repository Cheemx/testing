package main

import "sync"

// When to use Locks over channels and Goroutines?
//
// -> Use channels when passing ownership of data
// -> Use mutexes for managing state
//
// Don't be afraid to use a sync.Mutex if that fits your problem best.
// Go is pragmatic in letting you use the tools that solve your problem best
// and not forcing you into one style of code.

// go vet: can alert you to some subtle bugs in your code,
// USE IN Production to avoid small bugs

// Don't use embedding with muteses like following because it seems convenient
// it makes your mutex public to call and you'd want to be careful about
// APIs you're exposing!

// embedding example
// type Counter struct {
// 	sync.Mutex
// 	value int
// }

type Counter struct {
	// A Mutex is a mutual exclusion lock.
	// The zero value for a Mutex is an unlocked mutex.

	// A Mutex not be copied after first use.
	// That's why always pass pointers of structs containing mutexes
	mu    sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

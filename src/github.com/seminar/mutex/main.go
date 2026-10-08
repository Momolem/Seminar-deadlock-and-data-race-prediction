package main

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func main() {
	c := &Counter{}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		c.Increment()
	}()
	go func() {
		defer wg.Done()
		c.Increment()
	}()
	wg.Wait()
}

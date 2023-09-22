package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	wg := &sync.WaitGroup{}
	inc := Incrementer{}

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			for j := 0; j < 1000; j++ {
				inc.Increment()
			}

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(inc.GetI())
}

type Incrementer struct {
	i int64
}

func(inc *Incrementer) Increment() {
	atomic.AddInt64(&inc.i, 1)
}

func (inc *Incrementer) GetI() int64 {
	return inc.i
}
package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("REGULAR MAP:")
	RegularMap()
	fmt.Println("SYNC MAP:")
	SyncMap()
}

// Используем обычную Map
func RegularMap() {
	m := make(map[int]struct{})
	mut := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(num int) {
			for i := 0; i < 100; i++ {
				// Для использования Map из нескольких горутин требуется
				// оборачивать вызов map в Mutext
				mut.Lock()
				m[i + num] = struct{}{}
				mut.Unlock()
			}

			wg.Done()
		}(i)
	}

	wg.Wait()

	for k := range m {
		fmt.Printf("Key from regular map: %d\n", k)
	}
}

// Используем Sync map
func SyncMap() {
	m := sync.Map{}
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(num int) {
			for i := 0; i < 100; i++ {
				// Sync Map изначально потокобезопасна
				m.Store(i + num, struct{}{})
			}

			wg.Done()
		}(i)
	}

	wg.Wait()

	m.Range(func(key any, value any) bool{
		fmt.Printf("Key from sync map: %v\n", key)

		return true
	})
}
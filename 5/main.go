package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	wg.Add(1)

	// Остановка с помощью канала
	go func(ch chan struct{}) {
		for {
			select {
			case <-ch:
				fmt.Println("Gorutine 1 stopped...")
				wg.Done()
				return
			default:
				fmt.Println("Gorutine 1 steel working")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ch)

	// Остановка с помощью контекста
	wg.Add(1)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Gorutine 2 stopped...")
				wg.Done()
				return
			default:
				fmt.Println("Gorutine 2 steel working")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)

	// Остановка с помощью завершения основной горутины
	go func(ctx context.Context) {
		for {
			fmt.Println("Gorutine 3 steel working")
			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)

	// Ждем 5 секунд
	<- time.After(5 * time.Second)
	// Отменяем контекст
	cancel()
	// Отправляем пустую структуру в канал
	ch <- struct{}{}
	// ждем завершения всех горутин
	wg.Wait()
}

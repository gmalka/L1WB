package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// Время, которое будут записываться данные в канал
	n := 3
	ch := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		// Читаем данные из канала
		for v := range ch {
			fmt.Println(v)
		}
		wg.Done()
	}()

	// Канал, который отправляет значение через заданное время
	to := time.After(time.Duration(n) * time.Second)
	// Задаем семя для генерации случайных чисел
	rand.Seed(time.Now().Unix())
	for {
		// Select для определения того, закончился ли таймер или нет
		select {
		// Если таймер закончился то закрываем канал, и ждем завершения горутины
		case <- to:
			fmt.Println("Stopped")
			close(ch)
			wg.Wait()
			return
		// Если нет, то отправляем случайное значение в канал, и ждем небольшое время чтобы не спамить в терминал
		default:
			ch <- rand.Intn(1000)
			time.Sleep(300 * time.Millisecond)
		}
	}
}
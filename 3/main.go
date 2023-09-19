package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	DoIt(2)
}

// workCount - количество воркеров
func DoIt(workCount int) {
	// Канал куда будут отправляться сигналы
	sig := make(chan os.Signal, 1)
	// Регестрируем канал
	signal.Notify(sig, syscall.SIGINT)
	// Канал, куда будем отправлять данные на обработку и откуда
	// воркеры будут затем читать
	ch := make(chan int)

	wg := &sync.WaitGroup{}
	for i := 1; i <= workCount; i++ {
		wg.Add(1)
		go func(chanNum int) {
			// Читаем из канала пока он не закроется
			for v := range ch {
				fmt.Printf("Worker %d read %d\n", chanNum, v)
				time.Sleep(2 * time.Second)
			}

			wg.Done()
		}(i)
	}

	rand.Seed(time.Now().Unix())

	for {
		time.Sleep(500 * time.Millisecond)

		// Через select слушаем, нельзя ли считать данные с канала с сигналом
		select {
		// Если есть сигнал, то закрываем наш канал с данными, и ждем завершения воркеров
		case <- sig:
			close(ch)
			fmt.Println("server stop...")
			wg.Wait()
			return
		// если сигнала нет, то отправляем данные в канал откуда читают воркеры
		default:
			ch <- rand.Intn(1000)
		}
	}
}

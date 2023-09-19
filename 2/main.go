package main

import (
	"fmt"
	"sync"
)

func main() {
	First()
	Second()
}

// Первый Вариант
func First() {
	// исходный массив
	mas := []int{2, 4, 6, 8, 10}

	// WaitGroup для ожидания горутин
	wg := &sync.WaitGroup{}

	// Переменная, в которую будут складываться промежуточные значения
	sum := 0

	// Канал для получения промежуточных данных из горутин
	ch := make(chan int)

	for _, v := range mas {
		wg.Add(1)

		go func(num int) {
			// Записываем в канал квадрат числа
			ch <- (num * num)

			wg.Done()
		}(v)
	}

	// Еще один WaitGroup для ожидания чтения данных из канала
	sec := &sync.WaitGroup{}
	// Функция, которая читает данные из канала и прибавляет их к sum
	sec.Add(1)
	go func() {
		for v := range ch {
			sum += v
		}

		sec.Done()
	}()

	wg.Wait()

	// Закрываем канал после работы с ним
	close(ch)

	sec.Wait()

	fmt.Printf("Your sum is %d\n", sum)
}

// Второй Вариант
func Second() {
	// исходный массив
	mas := []int{2, 4, 6, 8, 10}

	// WaitGroup для ожидания горутин
	wg := &sync.WaitGroup{}

	// Переменная, в которую будут складываться промежуточные значения
	sum := 0

	// Мьютекс для одновременной записи в sum
	m := &sync.Mutex{}

	for _, v := range mas {
		wg.Add(1)

		go func(num int) {
			// Делаем Lock чтобы доступ к ресурсу был только у нашей горутины
			m.Lock()
			sum += num * num
			// Разблокируем ресурс
			m.Unlock()

			wg.Done()
		}(v)
	}

	wg.Wait()
	fmt.Printf("Your sum is %d\n", sum)
}
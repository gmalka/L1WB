package main

import "fmt"

func main() {
	// Начальное число
	var num int64 = 8

	// в какоое значение нужно установить i-й бит, 0 или 1
	byteToSet := int64(1)

	// на сколько битов нужно смместиться с 0-го
	i := int64(6)

	// самое смещение
	byteToSet = byteToSet << i

	// Обнуляем i-q бит
	num = num &^ (1 << i)

	// Заменяем i-й бит на наш
	num = num | byteToSet

	fmt.Println(num)

	for num > 0 {
		defer fmt.Print(num & 1)
		num = num >> 1
	}
}
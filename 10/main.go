package main

import "fmt"

func main() {
	// Слайсы для которых ищется пересечение
	first := []int{1, 5, 4, 8, 10, 55, 41, 0, 0}
	second := []int{6, 4, 1, 8, 9, -1, 2, 14, 68, 10}

	// Слайс, куда будет записываться результат
	result := []int{}

	m := make(map[int]struct{}, len(first))

	// Ключем мапы является значение первого слайся, заполняем ими нашу мапу
	for _, v := range first {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
		}
	}

	// Теперь проходимся по второму слайся и сверяем, есть ли в мапе уже значения по ключу,
	// который уже присутствует в мапе
	for _, v := range second {
		if _, ok := m[v]; ok {
			result = append(result, v)
		}
	}

	// Вывод результата
	fmt.Println(result)
}
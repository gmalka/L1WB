package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	m := make(map[string]struct{})

	// Записываем данные в мапу, добавляя в качестве ключа наши строки из слайса.
	// Так как ключи в мапе уникальны, то по итогу у нас окажется по одной копии каждого ключа
	for _, v := range arr {
		m[v] = struct{}{}
	}

	for k := range m {
		fmt.Println(k)
	}
}
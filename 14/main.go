package main

import (
	"math/rand"
	"time"
)

func createHugeString(count int) string {
	result := make([]byte, count)

	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < count; i++ {
		n := r.Intn(50)
		if n >= 25 {
			result[i] = byte(n - 25 + 97)
		} else {
			result[i] = byte(n + 65)
		}
	}

	return string(result)
}

var justString string
func someFunc() {
	// вместо создания двух строк, создаем одну
  justString = createHugeString(100)
}

func main() {
  someFunc()
}
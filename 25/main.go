package main

import "fmt"

func main() {
	str := "abCdefAaf"

	m := make(map[byte]struct{})

	for i := 0; i < len(str); i++ {
		s := str[i]
		if s >= 65 && s <= 90 {
			s += 32
		}

		if _, ok := m[s]; ok {
			fmt.Println(false)
			return
		}

		m[s] = struct{}{}
	}

	fmt.Println(true)
}
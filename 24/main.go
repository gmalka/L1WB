package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	Sleep(500 * time.Millisecond)

	fmt.Println(time.Now().Sub(now))
}

func Sleep(d time.Duration) {
	want := time.Now().Add(d)

	for time.Now().Before(want) {
	}

	return
}
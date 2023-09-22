package main

import (
	"fmt"
	"math/big"
)

func main() {
	first := &big.Int{}
	second := &big.Int{}

	first.SetString("20009182123124", 10)
	second.SetString("10000000000000", 10)

	t := big.Int{}
	t.Mul(first, second)
	fmt.Println("Умножение: ", t.String())

	t.Div(first, second)
	fmt.Println("Деление: ", t.String())

	t.Add(first, second)
	fmt.Println("Сложение: ", t.String())

	t.Sub(first, second)
	fmt.Println("Вычитание: ", t.String())
}
package main

import "fmt"


// Родительская структура Human с несколькими полями
type Human struct {
	name string
	age int
}

// Метод Introduce у структуры Human для того чтобы можно было представиться
func(h Human) Introduce() {
	fmt.Printf("Hi, my name is %s, i'am %d years old\n", h.name, h.age)
}

// Метод SayHelloTo для того чтобы Human мог поздороватсья с кем-нибуть
func(h Human) SayHelloTo(name string) {
	fmt.Printf("Hello %s from human\n", name)
}

// Дочерняя структура Action, которая содержит в себе анонимное поле Human, что позволяет
// работать с Action так, будто у него есть те же методы и поля что и у Human
type Action struct {
	Human
}

// Переопределение метода SayHelloTo, теперь если при работе с Action явно не вызвать метод у Human,
// то по умолчанию будет использоваться метод SayHelloTo из Action
func(a Action) SayHelloTo(name string) {
	fmt.Printf("Hello %s from action?WTF?!\n", name)
}

func main() {
	a := Action{Human: Human{name:"Anton", age:25}}

	// Introduce из Human
	a.Introduce()

	// Но SayHelloTo вызывается у Action
	a.SayHelloTo("Magomed")

	// Можем явно вызвать SayHelloTo н Human
	a.Human.SayHelloTo("Magomed")
}
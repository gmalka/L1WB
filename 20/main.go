package main

import "fmt"

// Наш интерфейс, который должны реализовывать наши объекты чтобы использоваться в дальнейшем
type HardWorder interface {
	HardWork() string
}

// Обычный человек, он соответствует нашему интерфейсу
type RegularHuman struct {
	Name string
	WorkCount int
}

// Тяжелая работа обычного человека
func (r RegularHuman) HardWork() string {
	return fmt.Sprintf("Человек %s работал %d часов и очень устал", r.Name, r.WorkCount)
}

// Антон. Он не соответствует интерфейсу HardWorker 
type AntonHuman struct {
	WorkCount int
}

// Антон не умеет работать усердно
func (r AntonHuman) Work() string {
	return fmt.Sprintf("Антон работал %d часов", r.WorkCount)
}

// Адаптер, чтобы превратить простую работу Антона в усердную работу
type AntonAdapter struct {
	anton AntonHuman
}

func (r AntonAdapter) HardWork() string {
	return fmt.Sprintf("%s и очень устал", r.anton.Work())
}

func main() {
	var worker HardWorder

	worker = RegularHuman{
		Name: "Витя",
		WorkCount: 18,
	}
	fmt.Println(worker.HardWork())

	worker = AntonAdapter{
		anton: AntonHuman{
			WorkCount: 10,
		},
	}

	fmt.Println(worker.HardWork())
}
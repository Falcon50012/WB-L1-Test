package main

import "fmt"

// Структура Human: Содержит поле Name типа string и поле Action типа Action, являющееся встроенной структурой.
type Human struct {
	Name   string
	Action Action
}

// Метод Intro(): Метод структуры Human.
func (h Human) Intro() {
	h = Human{
		"Thomas Anderson",
		Action{Enter: "Outside the Matrix"},
	}

	fmt.Println(h)
}

// Структура Action.
type Action struct {
	Enter string
}

// Метод структуры Action.
func (a Action) EnterTheMatrix() {
	h := Human{
		"Neo",
		Action{Enter: "Inside the Matrix"},
	}

	fmt.Println(h)
}

func main() {
	h := Human{}
	h.Intro()
	a := Action{}
	a.EnterTheMatrix()
}

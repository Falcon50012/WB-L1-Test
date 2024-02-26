package main

import "fmt"

// Структура Human: Содержит поле Name типа string и поле Action типа *Action.
type Human struct {
	Name   string
	Action *Action
}

// Метод Intro(): Метод структуры Human.
func (h *Human) Intro() {
	h.Name = "Thomas Anderson"
	h.Action = &Action{Enter: "Outside the Matrix"}
	fmt.Println(h.Name, *h.Action) // Разыменовываем указатель Action
}

// Структура Action.
type Action struct {
	Enter string
}

// Метод структуры Action.
func (a *Action) EnterTheMatrix() {
	h := &Human{
		Name:   "Neo",
		Action: &Action{Enter: "Inside the Matrix"},
	}
	fmt.Println(h.Name, *h.Action)
}

func main() {
	h := Human{}
	h.Intro()
	a := Action{}
	a.EnterTheMatrix()
}

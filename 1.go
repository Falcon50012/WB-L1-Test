package main

import "fmt"

// Структура Human: Содержит поле Name типа string и поле Action типа Action, являющееся вложенной структурой.
type Human struct {
	Name   string
	Action *Action
}

// Метод Intro(): Метод структуры Human.
func (h *Human) Intro() {
	h.Name = "Thomas Anderson"
	h.Action = &Action{"Outside the Matrix"}
	fmt.Println(h.Name, h.Action.Enter)
}

// Структура Action.
type Action struct {
	Enter string
}

// Метод структуры Action.
func (a *Action) EnterTheMatrix(h *Human) {
	h.Name = "Neo"
	a.Enter = "Inside the Matrix"
	fmt.Println(h.Name, a.Enter)
}

func main() {
	h := Human{}
	h.Intro()
	h.Action.EnterTheMatrix(&h)
}

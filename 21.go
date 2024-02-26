//Реализовать паттерн «адаптер» на любом примере.

package main

import "fmt"

// Интерфейс геймпада PlayStation
type PlayStationGamepad interface {
	ConnectToPlayStation() string
}

// Имплементация геймпада PlayStation
type DualShock4 struct{}

func (d *DualShock4) ConnectToPlayStation() string {
	return "DualShock 4 connected"
}

// Интерфейс ПК-геймпада
type PCGamepad interface {
	ConnectToPC() string
}

// Имплементация ПК-геймпада
type XInputGamepad struct{}

func (x *XInputGamepad) ConnectToPC() string {
	return "XInput Gamepad connected to PC"
}

// Адаптер для подключения геймпада PlayStation к ПК
type PlayStationToPCAdapter struct {
	playStationGamepad PlayStationGamepad
}

func (a *PlayStationToPCAdapter) ConnectToPC() string {
	// Преобразуем вызов метода геймпада PlayStation в метод ПК-геймпада
	psConnectMsg := a.playStationGamepad.ConnectToPlayStation()
	return fmt.Sprintf("Adapter: %s via adapter to PC", psConnectMsg)
}

func main() {
	// Создаем геймпад PlayStation и адаптер
	playStationGamepad := &DualShock4{}
	adapter := &PlayStationToPCAdapter{playStationGamepad}

	// Подключаем геймпад PlayStation к ПК с помощью адаптера
	fmt.Println(adapter.ConnectToPC())
}

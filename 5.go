//Разработать программу, которая будет последовательно отправлять значения в канал,
//а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

package main

import (
	"fmt"
	"time"
)

func sender(ch chan<- int) {
	for i := 1; ; i++ {
		ch <- i
		time.Sleep(time.Second) // Отправляем значение в канал каждую секунду
	}
}

func main() {
	var countdown int
	fmt.Println("Введите значение таймера: ")
	fmt.Scan(&countdown)

	if countdown <= 0 {
		fmt.Println("Завершение программы...")
		return
	}

	ch := make(chan int)

	// Запускаем горутину, которая отправляет значения в канал
	go sender(ch)

	// Устанавливаем таймер на N секунд
	duration := time.Duration(countdown) * time.Second
	timer := time.NewTimer(duration)

	// Читаем значения из канала, пока работает таймер
	for {
		select {
		case <-timer.C: //field C <-chan Time of Timer type
			fmt.Println("Программа завершена")
			return
		case val := <-ch:
			fmt.Println("Принято значение из канала:", val)
		}
	}
}

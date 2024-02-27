//Реализовать собственную функцию sleep.

package main

import (
	"fmt"
	"time"
)

// Функция sleep, которая останавливает выполнение программы на заданное количество секунд
func sleep(seconds int) {
	<-time.After(time.Duration(seconds) * time.Second)
}

func main() {
	fmt.Println("Начало программы")
	sleep(3) // Приостанавливаем выполнение на 3 секунды
	fmt.Println("Прошло 3 секунды")
	fmt.Println("Завершение программы")
}

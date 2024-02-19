//Разработать конвейер чисел.
//Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
//после чего данные из второго канала должны выводиться в stdout.

package main

import (
	"fmt"
)

// Этап 1: Генерация целых чисел в диапазоне от start до end и отправка их в канал
func generateNumbers(start, end int, out chan<- int) {
	defer close(out)
	for i := start; i <= end; i++ {
		out <- i
	}
}

// Этап 2: Возведение чисел в квадрат и отправка результатов в канал
func squareNumbers(in <-chan int, out chan<- int) {
	defer close(out)
	for num := range in {
		out <- num * num
	}
}

// Этап 3: Вывод чисел в консоль
func printNumbers(in <-chan int) {
	for num := range in {
		fmt.Println(num)
	}
}

func main() {
	// Создаем каналы для связи между этапами конвейера
	numbers := make(chan int)
	squares := make(chan int)

	// Запускаем горутину для генерации чисел
	go generateNumbers(1, 5, numbers)

	// Запускаем горутину для возведения чисел в квадрат
	go squareNumbers(numbers, squares)

	// Выводим результаты
	printNumbers(squares)
}

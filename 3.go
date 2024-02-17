package main

import (
	"fmt"
)

func squareWorker3(num int, ch1 chan<- int) {
	ch1 <- num * num
}

func squaredAddWorker(squared int, ch2 chan<- int) {
	ch2 <- squared + squared
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	ch1 := make(chan int, len(numbers)) // буферизованный канал
	ch2 := make(chan int, len(numbers)) // буферизованный канал

	// Запускаем горутины для вычисления квадратов чисел
	for _, num := range numbers {
		go squareWorker3(num, ch1)
	}

	// Получаем квадраты чисел и суммируем их
	fmt.Print("Квадраты чисел: ")
	for range numbers {
		squared := <-ch1
		fmt.Print(squared, " ")
		go squaredAddWorker(squared, ch2)
	}

	//fmt.Println()

	// Получаем сумму квадратов чисел
	sum := 0
	for range numbers {
		squaredAdd := <-ch2
		//fmt.Print(addSquared, " ")
		sum += squaredAdd
	}

	// Выводим результат
	fmt.Println("\nСумма квадратов чисел:", sum)
}

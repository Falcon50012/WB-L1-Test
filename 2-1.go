//Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10)
//и выведет их квадраты в stdout

package main

import "fmt"

func squareWorker(num int, ch chan<- int) {
	ch <- num * num
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	numWorkers := len(numbers)
	ch := make(chan int, numWorkers) // Создание буферизированного канала

	for _, num := range numbers {
		go squareWorker(num, ch)
	}

	for range numbers {
		squared := <-ch
		fmt.Println(squared)
	}
}

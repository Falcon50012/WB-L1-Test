package main

import "fmt"

func squareWorker2(num int, ch chan<- int) {
	ch <- num * num
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	numWorkers := len(numbers)
	ch := make(chan int, numWorkers) // Создание буферезированного канала

	for _, num := range numbers {
		go squareWorker2(num, ch)
	}

	for range numbers {
		squared := <-ch
		fmt.Println(squared)
	}
}

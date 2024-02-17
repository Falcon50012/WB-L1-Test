package main

import "fmt"

func squareWorker(num int, ch chan<- int) {
	ch <- num * num
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	ch := make(chan int) //Создание небуферизированного канала

	for _, num := range numbers {
		go squareWorker(num, ch)
	}

	for range numbers {
		squared := <-ch
		fmt.Println(squared)
	}
}

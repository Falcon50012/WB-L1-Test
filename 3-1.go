package main

import (
	"fmt"
	"sync"
)

func squareWorker4(num int, ch1 chan<- int) {
	ch1 <- num * num
}

func addSquareWorker(sqr int, ch2 chan<- int) {
	ch2 <- sqr + sqr
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	ch1 := make(chan int, len(numbers))
	ch2 := make(chan int, len(numbers))
	var wg sync.WaitGroup

	// Запускаем горутины для вычисления квадратов чисел
	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			squareWorker4(n, ch1)
		}(num)
	}

	// После вывода squared - fatal error: all goroutines are asleep - deadlock!
	//for range numbers {
	//	squared := <-ch1
	//	fmt.Println(squared)
	//}

	// Получаем квадраты чисел и суммируем их
	go func() {
		wg.Wait()
		close(ch1)
	}()

	for squared := range ch1 {
		go addSquareWorker(squared, ch2)
	}

	// Получаем сумму квадратов чисел
	sum := 0
	for range numbers {
		addSquared := <-ch2
		sum += addSquared
	}

	// Выводим результат
	fmt.Println("Сумма квадратов чисел:", sum)
}

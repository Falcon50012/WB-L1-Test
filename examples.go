//Применение каналов для сигнализации о завершении работы горутин:
//В этом подходе используется дополнительный канал для сигнализации о завершении работы горутин.
//Когда все горутины завершают свою работу, они отправляют сигнал в этот канал.
//Вот пример:

package main

import (
	"fmt"
	"sync"
)

func worker(id int, tasks <-chan int, results chan<- int, done chan<- struct{}) {
	defer func() {
		done <- struct{}{}
	}()
	for task := range tasks {
		results <- task * task
	}
}

func main() {
	numWorkers := 4
	tasks := make(chan int, 10)
	results := make(chan int, 10)
	done := make(chan struct{}, numWorkers)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, tasks, results, done)
		}(i)
	}

	// Заполнение канала задач
	for i := 0; i < 10; i++ {
		tasks <- i
	}
	close(tasks)

	// Ожидание завершения работы всех горутин
	go func() {
		wg.Wait()
		close(results)
	}()

	// Ожидание завершения работы всех горутин
	for i := 0; i < numWorkers; i++ {
		<-done
	}

	// Получение результатов
	for result := range results {
		fmt.Println(result)
	}
}

//Использование пакета sync/atomic для безопасного доступа к общим данным:
//В этом подходе используется пакет sync/atomic для безопасной работы с общими данными из нескольких горутин.
//Вот пример:

//package main
//
//import (
//"fmt"
//"sync"
//"sync/atomic"
//)
//
//func main() {
//	var sum int64
//	numWorkers := 4
//	var wg sync.WaitGroup
//
//	for i := 0; i < numWorkers; i++ {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			for i := 0; i < 1000; i++ {
//				atomic.AddInt64(&sum, int64(i))
//			}
//		}()
//	}
//
//	wg.Wait()
//
//	fmt.Println("Sum:", sum)
//}

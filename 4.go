package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func publisher(ch chan<- int) {
	rand.Seed(time.Now().UnixNano())
	for {
		randomBinary := rand.Intn(2) // Генерация случайного бинарного числа (0 или 1)
		ch <- randomBinary
		time.Sleep(400 * time.Millisecond)
	}
}

func worker2(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range ch {
		fmt.Printf("Worker %d: Received data: %d\n", id, data)
		// Здесь можно добавить обработку полученных данных
	}
}

func main() {
	var workers int = 8
	//fmt.Println("Выберите количесиво воркеров")
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(workers)

	for i := 1; i <= workers; i++ {
		go worker2(i, ch, &wg)
	}

	go publisher(ch)

	// Ожидание завершения всех воркеров
	wg.Wait()
}

//package main
//
//import (
//"context"
//"fmt"
//"math/rand"
//"os"
//"os/signal"
//"sync"
//"syscall"
//"time"
//)
//
//func publisher(ctx context.Context, ch chan<- int) {
//	rand.Seed(time.Now().UnixNano())
//	for {
//		select {
//		case <-ctx.Done():
//			return // Завершаем работу при получении сигнала об отмене
//		default:
//			randomBinary := rand.Intn(2) // Генерация случайного бинарного числа (0 или 1)
//			ch <- randomBinary
//			time.Sleep(400 * time.Millisecond)
//		}
//	}
//}
//
//func worker(ctx context.Context, id int, ch <-chan int, wg *sync.WaitGroup) {
//	defer wg.Done()
//	for {
//		select {
//		case <-ctx.Done():
//			return // Завершаем работу при получении сигнала об отмене
//		case data := <-ch:
//			fmt.Printf("Worker %d: Received data: %d\n", id, data)
//			// Здесь можно добавить обработку полученных данных
//		}
//	}
//}
//
//func main() {
//	var workers int = 8
//	ch := make(chan int)
//	var wg sync.WaitGroup
//	wg.Add(workers)
//
//	// Создаем контекст для отмены работы воркеров
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel() // Отменяем контекст при завершении работы программы
//
//	// Запуск воркеров
//	for i := 0; i < workers; i++ {
//		go worker(ctx, i, ch, &wg)
//	}
//
//	// Запуск производителя данных
//	go publisher(ctx, ch)
//
//	// Обработка сигнала Ctrl+C для отмены работы программы
//	sigCh := make(chan os.Signal, 1)
//	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
//	<-sigCh
//
//	// Отмена работы всех воркеров
//	cancel()
//
//	// Ожидание завершения всех воркеров
//	wg.Wait()
//}

package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func publisher(ch chan<- int) {
	rand.Seed(time.Now().UnixNano())
	for {
		randomBinary := rand.Intn(10)
		ch <- randomBinary
		time.Sleep(400 * time.Millisecond)
	}
}

func worker(id int, ch <-chan int, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Воркер %d завершает работу...\n", id)
			return
		case data := <-ch:
			fmt.Printf("Воркер %d: Принял данные: %d\n", id, data)
		}
	}
}

func main() {
	var workers int
start:
	fmt.Print("Выберите количество воркеров: ")
	fmt.Scan(&workers)
	if workers <= 0 {
		fmt.Println("Количество воркеров должно быть >= 0!")
		goto start
	}

	ch := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(workers)
	for i := 1; i <= workers; i++ {
		go worker(i, ch, &wg, ctx)
	}

	go publisher(ch)

	// Обработка сигнала Ctrl+C
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh
	fmt.Println("\nПрограмма завершается...")

	// Отмена работы всех воркеров
	cancel()

	// Ожидание завершения всех воркеров
	wg.Wait()
}

//Реализовать все возможные способы остановки выполнения горутины.
//1. Использование канала для сигнала остановки

package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan struct{})
	go func() {
		for {
			select {
			default:
				fmt.Println("Working...")
				time.Sleep(1 * time.Second)
			case <-stop:
				fmt.Println("Stopping...")
				return
			}
		}
	}()

	time.Sleep(3 * time.Second)
	close(stop)
	time.Sleep(1 * time.Second)
	fmt.Println("Main goroutine exits")
}

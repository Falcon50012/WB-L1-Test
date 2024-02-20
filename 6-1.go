//Реализовать все возможные способы остановки выполнения горутины.
//2. Использование контекста

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			default:
				fmt.Println("Working...")
				time.Sleep(time.Second)
			case <-ctx.Done():
				fmt.Println("Stopping...")
				return
			}
		}
	}(ctx)

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
	fmt.Println("Main goroutine exits")
}

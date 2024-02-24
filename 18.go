//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
)

// Структура-счетчик
type Counter struct {
	mu    sync.Mutex
	value int
}

// Метод для инкрементирования счетчика
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Метод для получения текущего значения счетчика
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	// Создаем экземпляр счетчика
	counter := Counter{}

	// Количество горутин для инкрементирования счетчика
	numWorkers := 10

	// Создаем горутины для инкрементирования счетчика
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Выводим итоговое значение счетчика
	fmt.Println("Итоговое значение счетчика:", counter.Value())
}

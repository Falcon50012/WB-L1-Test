//Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	workers := 11
	m := make(map[int]float64)

	// Функция для конкурентной записи в map
	writeToMap := func(key int, value float64) {
		defer wg.Done()
		mu.Lock()
		defer mu.Unlock()
		m[key] = value
	}

	// Запуск нескольких горутин для записи в map
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go writeToMap(i, math.Pow(2, float64(i)))
	}

	wg.Wait()

	for i, value := range m {
		fmt.Printf("Значение для ключа %d: %d\n", i, int64(value))
	}

	fmt.Println(m)

	fmt.Println("Главная горутина завершена")
}

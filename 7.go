//Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	workers := 11
	m := make(map[int]float64)

	// Функция для конкурентной записи в map
	writeToMap := func(key int, value float64) {
		mu.Lock()
		defer mu.Unlock()
		m[key] = value
	}

	// Запуск нескольких горутин для записи в map
	for i := 0; i < workers; i++ {
		go writeToMap(i, math.Pow(2, float64(i)))
	}

	//Ожидание завершения всех горутин
	for i := 0; i < workers; i++ {
		go func(i int) {
			mu.Lock()
			defer mu.Unlock()
			fmt.Printf("Значение для ключа %d: %d\n", i, int64(m[i]))
			//fmt.Printf("Горутина %v завершена\n", i)
		}(i)
	}

	time.Sleep(5 * time.Second)
	fmt.Println(m)
	fmt.Println("Главная горутина завершена")
}

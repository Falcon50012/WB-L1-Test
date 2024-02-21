//Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
//Объединить данные значения в группы с шагом в 10 градусов.
//Последовательность в подмножноствах не важна.

package main

import (
	"fmt"
	"math"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float64)

	for _, temp := range temperatures {
		// Определяем ключ диапазона
		key := int(math.Floor(temp/10)) * 10
		// Добавляем температуру в соответствующую группу
		groups[key] = append(groups[key], temp)
	}

	// Выводим результат
	for key, group := range groups {
		fmt.Printf("Диапазон %d - %d: %v\n", key, key+10, group)
	}
}

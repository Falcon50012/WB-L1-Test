//Реализовать бинарный поиск встроенными методами языка.

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	var arr []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(10)+1)
	}

	// Исходный массив
	fmt.Println("Исходный массив", arr)

	// Сортировка массива
	sort.Ints(arr)

	// Вывод отсортированного массива
	fmt.Println("Отсортированный массив:", arr)

	// Искомое значение
	target := 10

	// Ищем индекс элемента в массиве
	index := sort.Search(len(arr), func(i int) bool { return arr[i] >= target })

	// Проверяем, найден ли элемент
	if index < len(arr) && arr[index] == target {
		fmt.Printf("Элемент %d найден на позиции %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден\n", target)
	}
}

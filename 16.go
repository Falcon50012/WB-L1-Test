//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

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
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	// Вывод отсортированного массива
	fmt.Println("Отсортированный массив:", arr)
}

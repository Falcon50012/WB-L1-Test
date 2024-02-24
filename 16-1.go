package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var arr []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(10)+1)
	}
	fmt.Println("Исходный массив:", arr)

	// Применяем быструю сортировку (quicksort)
	quicksort(arr, 0, len(arr)-1)

	fmt.Println("Отсортированный массив:", arr)
}

func quicksort(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		quicksort(arr, low, pivot-1)
		quicksort(arr, pivot+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

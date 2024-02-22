//package main
//
//import "fmt"
//
//func setBit(n int64, i int) int64 {
//	return n | (1 << i)
//}
//func clearBit(n int64, i int) int64 {
//	return n &^ (1 << i)
//}
//func main() {
//	num := int64(-10) // пример числа
//	fmt.Printf("Число: %d\n", num)
//	i := 5 // индекс бита для установки/очистки
//	// Установка бита
//	newNum := setBit(num, i)
//	fmt.Printf("После установки бита: %d (бит %d установлен)\n", newNum, i)
//
//	// Очистка бита
//	clearNum := clearBit(num, i)
//	fmt.Printf("После очистки бита: %d (бит %d очищен)\n", clearNum, i)
//}

//import (
//	"fmt"
//	"math"
//	"strconv"
//)
//
//func main() {
//	m := make(map[int]float64)
//	m[1] = 1.2
//
//	x := math.Pow(100, 0.5)
//	_ = x
//	x1 := math.Sqrt(100)
//	_ = x1
//
//	var num int64 = 10
//	//num << 1
//	bnum := strconv.FormatInt(num, 2)
//	fmt.Println(bnum)
//}

//func bubbleSort(arr []int) {
//	n := len(arr)
//	for i := 0; i < n-1; i++ {
//		for j := 0; j < n-i-1; j++ {
//			if arr[j] > arr[j+1] {
//				arr[j], arr[j+1] = arr[j+1], arr[j]
//			}
//		}
//	}
//}
//func main() {
//	rand.Seed(time.Now().UnixNano())
//	maxSize := 20
//	elements := 10
//	// Создаем массив случайных чисел
//	arr := make([]int, elements)
//	for i := range arr {
//		arr[i] = rand.Intn(maxSize)
//	}
//
//	fmt.Println("Original array:")
//
//	for _, v := range arr {
//		fmt.Print(v)
//		fmt.Print(" ")
//	}
//	fmt.Println()
//	bubbleSort(arr)
//
//	fmt.Println("Sorted array:")
//	for _, v := range arr {
//		fmt.Print(v)
//		fmt.Print(" ")
//	}
//	fmt.Println()
//}

//package main
//
//import (
//	"fmt"
//	"sort"
//)
//
//func main() {
//	numbers := []int{4, 2, 7, 1, 9, 5}
//
//	// Сортировка среза целых чисел в порядке возрастания
//	sort.Slice(numbers, func(i, j int) bool {
//		return numbers[i] < numbers[j]
//	})
//
//	fmt.Println("Отсортированный срез:", numbers)
//}

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

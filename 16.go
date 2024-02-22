//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	//var low, high int
	var arr []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(10) + 1
		//low = arr[0]
		//high = arr[i]
	}
	fmt.Println(arr)
	//fmt.Println(low)
	//fmt.Println(high)
	//mid := (low + high) / 2
	//fmt.Println(mid)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	fmt.Println(arr)
}

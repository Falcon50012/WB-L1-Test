//Удалить i-ый элемент из слайса.

package main

import "fmt"

func removeIndex(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5}

	// Удаляем 2-й элемент
	slice = removeIndex(slice, 1)

	fmt.Println(slice)
}

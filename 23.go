//Удалить i-ый элемент из слайса.

package main

import "fmt"

func removeIndex(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	// Пример слайса
	slice := []int{1, 2, 3, 4, 5}

	// Удаляем второй элемент (индекс 1)
	slice = removeIndex(slice, 1)

	// Выводим результат
	fmt.Println(slice) // [1 3 4 5]
}

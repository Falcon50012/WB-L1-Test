//Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

package main

import (
	"fmt"
)

// Функция, которая переворачивает строку
func reverseString(s string) string {
	// Преобразуем строку в массив рун
	runes := []rune(s)
	// Получаем длину строки в рунах
	length := len(runes)

	// Переворачиваем руны в массиве
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Возвращаем перевернутую строку
	return string(runes)
}

func main() {
	// Исходная строка
	input := "главрыба"
	fmt.Println("Исходная строка:", input)

	// Переворачиваем строку
	reversed := reverseString(input)
	fmt.Println("Перевернутая строка:", reversed)
}

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

package main

import (
	"fmt"
	"strings"
)

// Функция, которая переворачивает слова в строке
func reverseWords(s string) string {
	// Разделяем строку на слова
	words := strings.Fields(s)
	// Переворачиваем порядок слов
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	// Объединяем слова в строку с пробелами между ними
	return strings.Join(words, " ")
}

func main() {
	// Исходная строка
	input := "snow dog sun"
	fmt.Println("Исходная строка:", input)

	// Переворачиваем слова в строке
	reversed := reverseWords(input)
	fmt.Println("Строка с перевернутыми словами:", reversed)
}

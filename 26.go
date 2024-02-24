//Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
//Функция проверки должна быть регистронезависимой.

package main

import (
	"fmt"
	"strings"
)

// Функция, которая проверяет, что все символы в строке уникальные, игнорируя регистр
func isUnique(str string) bool {
	// Преобразуем строку к нижнему регистру
	str = strings.ToLower(str)

	// Создаем карту для подсчета количества встреченных символов
	charCount := make(map[rune]int)

	// Проходим по каждому символу в строке
	for _, char := range str {
		// Если символ уже встречался, значит строка не уникальна
		if charCount[char] > 0 {
			return false
		}
		// Увеличиваем счетчик встреченного символа
		charCount[char]++
	}

	// Если цикл завершился без прерывания, значит все символы уникальны
	return true
}

func main() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"

	fmt.Printf("%s: %t\n", str1, isUnique(str1))
	fmt.Printf("%s: %t\n", str2, isUnique(str2))
	fmt.Printf("%s: %t\n", str3, isUnique(str3))
}

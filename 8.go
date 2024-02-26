//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// setBitToValue устанавливает i-й бит в num в bitValue (0 или 1)
func setBitToValue(num int64, pos uint, bitValue int64) int64 {
	mask := int64(1 << pos) // Создаем маску с единицей в позиции i

	if bitValue == 1 {
		return num | mask // Устанавливаем i-й бит в 1 с помощью побитовой операции ИЛИ
	}
	return num &^ mask // Устанавливаем i-й бит в 0 с помощью побитовой операции AND с инвертированной маской
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Создаем случайное число
	//num := rand.Int63n(100)
	var num int64 = 10
	binNum := strconv.FormatInt(num, 2)
	// Выводим исходное число
	fmt.Printf("Исходное число: %d\n", num)
	fmt.Println("binNum:", binNum)
	// Устанавливаем i-й бит в 1
	i := uint(5)         // Пример: установить 5-й бит
	bitValue := int64(1) // Установить в 1
	num = setBitToValue(num, i, bitValue)
	fmt.Printf("Число после установки %d-го бита в 1: %d\n", i, num)

	// Устанавливаем i-й бит в 0
	bitValue = 0 // Установить в 0
	num = setBitToValue(num, i, bitValue)
	fmt.Printf("Число после установки %d-го бита в 0: %d\n", i, num)
}

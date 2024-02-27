//Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Создаем объекты big.Int для a и b, которые больше чем 2^20
	a := new(big.Int).Exp(big.NewInt(2), big.NewInt(22), nil)
	b := new(big.Int).Exp(big.NewInt(2), big.NewInt(21), nil)

	// Умножение
	multiplicationResult := new(big.Int).Mul(a, b)
	fmt.Printf("Умножение: %s * %s = %s\n", a, b, multiplicationResult)

	// Деление
	divisionResult := new(big.Int).Div(a, b)
	fmt.Printf("Деление: %s / %s = %s\n", a, b, divisionResult)

	// Сложение
	additionResult := new(big.Int).Add(a, b)
	fmt.Printf("Сложение: %s + %s = %s\n", a, b, additionResult)

	// Вычитание
	subtractionResult := new(big.Int).Sub(a, b)
	fmt.Printf("Вычитание: %s - %s = %s\n", a, b, subtractionResult)
}

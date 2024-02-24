//Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Создаем объекты big.Int для a и b, которые больше чем 2^20
	a := new(big.Int).Exp(big.NewInt(2), big.NewInt(21), nil)
	b := new(big.Int).Exp(big.NewInt(2), big.NewInt(20), nil)

	// Умножение
	multiplicationResult := new(big.Int).Mul(a, b)
	fmt.Printf("Умножение: %s * %s = %s\n", a.String(), b.String(), multiplicationResult.String())

	// Деление
	divisionResult := new(big.Int).Div(a, b)
	fmt.Printf("Деление: %s / %s = %s\n", a.String(), b.String(), divisionResult.String())

	// Сложение
	additionResult := new(big.Int).Add(a, b)
	fmt.Printf("Сложение: %s + %s = %s\n", a.String(), b.String(), additionResult.String())

	// Вычитание
	subtractionResult := new(big.Int).Sub(a, b)
	fmt.Printf("Вычитание: %s - %s = %s\n", a.String(), b.String(), subtractionResult.String())
}

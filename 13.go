//Поменять местами два числа без создания временной переменной.

package main

import "fmt"

func main() {
	a, b := 5, 10

	fmt.Println("До обмена:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	a, b = b, a

	fmt.Println("После обмена:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}

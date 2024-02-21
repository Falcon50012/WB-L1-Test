package main

import (
	"fmt"
	"math"
)

//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func main() {
	var num int64 = -10
	//fmt.Println(num)

	bits := make([]int, 64)

	//В этом коде мы используем побитовый сдвиг вправо (>>) и побитовое И (&) для получения каждого бита числа num.
	//Затем мы записываем каждый бит в соответствующую ячейку массива.
	//Мы сдвигаем число num на i битов вправо с помощью оператора >>.
	//Это позволяет нам получить значение каждого бита по очереди.
	//Затем мы применяем побитовое И с числом 1, чтобы оставить только младший бит, который и записываем в переменную bit.
	for i := len(bits) - 1; i >= 0; i-- {
		bit := (num >> uint(i)) & 1
		//fmt.Print(bit)
		bits[i] = int(bit)
	}
	//fmt.Println()
	fmt.Printf("Исходное число в бинарном представлении:\n %v\n", bits)
	//fmt.Println(bits[6])

	//bitsValue := make([]uint64, 64)
	//var bitValue float64
	//
	//for i := 0; i < len(bits); i++ {
	//	bitValue = math.Pow(2, float64(i))
	//	intBitValue := uint64(bitValue)
	//	bitsValue[i] = intBitValue
	//}
	//fmt.Printf("Степени двойки в 64-битном числе:\n %v\n", bitsValue)

	//rand.Seed(time.Now().UnixNano())
	//for i := len(bits) - 1; i >= 0; i-- {
	//	bits[i] = rand.Intn(2)
	//}
	//fmt.Printf("Рандомное число в бинарном представлении:\n %v\n", bits)

	var newNum float64

	for i := 0; i < len(bits); i++ {
		if bits[i] == 1 {
			newNum += math.Pow(2, float64(i))
		}
	}
	fmt.Printf("Рандомное число в десятичном представлении:\n %v\n", int64(newNum))
}

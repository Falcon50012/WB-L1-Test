//К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
//Приведите корректный пример реализации.
//
//var justString string
//func someFunc() {
//  v := createHugeString(1 << 10)
//  justString = v[:100]
//}
//
//func main() {
//  someFunc()
//}
//
//Этот фрагмент кода может привести к проблемам с утечкой памяти из-за использования большой строки в
//createHugeString и сохранения только ее первых 100 символов в justString.
//Поскольку только первые 100 символов копируются в justString, а сама большая строка не освобождается,
//это может привести к накоплению неиспользуемой памяти.
//
//Чтобы исправить это, необходимо очистить ресурсы, выделенные для большой строки после использования.
//Мы можем сделать это, принудительно уничтожив большую строку после того, как нам больше не нужны ее данные.

package main

import (
	"fmt"
	"strconv"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	defer func() {
		// Очищаем ресурсы, выделенные для большой строки
		v = ""
	}()
	if len(v) >= 100 {
		justString = v[:100]
	} else {
		// Обработка случая, когда длина строки меньше 100
		justString = v
	}
}

func main() {
	someFunc()
	fmt.Println(justString)
}

func createHugeString(size int) string {
	// Преобразуем число в строку
	strSize := strconv.Itoa(size)
	return "Huge string of size " + strSize
}

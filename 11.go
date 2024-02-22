//Реализовать пересечение двух неупорядоченных множеств

package main

import "fmt"

// Функция для поиска пересечения двух множеств
func intersection(set1, set2 map[int]bool) map[int]bool {
	result := make(map[int]bool) //Ключи - элементы множества, значения - булевы флаги

	//Производим итерацию по элементам пеового множества
	for k := range set1 {
		//Если текущий эл-т присутствует во 2-м множестве, добавляем его в результат
		if set2[k] {
			result[k] = true
		}
	}

	return result
}

func main() {
	set1 := map[int]bool{1: true, 6: true, 3: true, 4: true, 8: true}
	set2 := map[int]bool{2: true, 3: true, 6: true, 10: true, 8: true}

	intersect := intersection(set1, set2)

	fmt.Println("Пересечение множеств:", intersect)
}

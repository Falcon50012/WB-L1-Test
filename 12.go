//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

package main

import "fmt"

func main() {
	strSequence := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]bool)

	for _, item := range strSequence {
		set[item] = true
	}

	fmt.Println("Множество", set)
}

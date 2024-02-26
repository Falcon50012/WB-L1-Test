//Разработать программу,
//которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

package main

import "fmt"

func getType(x interface{}) string {
	switch x.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return "channel of int"
	default:
		return "unknown"
	}
}

func main() {
	var a interface{} = 5
	var b interface{} = "hello"
	var c interface{} = true
	var d interface{} = make(chan int)

	fmt.Println("Type of a:", getType(a))
	fmt.Println("Type of b:", getType(b))
	fmt.Println("Type of c:", getType(c))
	fmt.Println("Type of d:", getType(d))
}

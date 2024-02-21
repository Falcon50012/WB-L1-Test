package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	m := make(map[int]float64)
	m[1] = 1.2

	x := math.Pow(100, 0.5)
	_ = x
	x1 := math.Sqrt(100)
	_ = x1

	var num int64 = 10
	//num << 1
	bnum := strconv.FormatInt(num, 2)
	fmt.Println(bnum)
}

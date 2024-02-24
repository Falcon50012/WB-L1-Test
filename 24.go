//Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point
//с инкапсулированными параметрами x,y и конструктором.

package main

import (
	"fmt"
	"math"
)

// Определение структуры Point
type Point struct {
	x, y float64
}

// Конструктор для создания новой точки
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// Метод для нахождения расстояния между текущей точкой и другой точкой
func (p Point) DistanceTo(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// Создание двух точек
	point1 := NewPoint(0, 0)
	point2 := NewPoint(3, 4)

	// Вычисление расстояния между точками
	distance := point1.DistanceTo(point2)

	// Вывод результата
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}

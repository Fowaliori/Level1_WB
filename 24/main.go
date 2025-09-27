package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками на плоскости.

Точки представлены в виде структуры Point с инкапсулированными (приватными) полями x, y (типа float64) и конструктором.

Расстояние рассчитывается по формуле между координатами двух точек.

Подсказка: используйте функцию-конструктор NewPoint(x, y), Point и метод Distance(other Point) float64.
*/

type Point struct {
	x, y float64
}

func (p *Point) NewPoint(x, y float64) {
	p.x = x
	p.y = y
}

func (p Point) Distance(p2 Point) float64 {
	d := math.Sqrt(math.Pow(p2.x-p.x, 2) + math.Pow(p2.y-p.y, 2))
	return d
}

func main() {
	var p1, p2 Point
	p1.NewPoint(4, 2)
	p2.NewPoint(3, 5)
	fmt.Println(p1.Distance(p2))
}

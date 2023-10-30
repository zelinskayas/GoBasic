package main

import (
	"fmt"
	"math"
)

func main() {
	var a, d, o float64
	fmt.Println("Ожидается ввод площади круга: ")
	fmt.Scan(&a)

	d = 2 * math.Sqrt(a/math.Pi)
	o = math.Pi * d

	fmt.Println("Диаметр окружности равен: ", d)
	fmt.Println("Длина окружности равна: ", o)
}

//d = 2 * √S : Пи.

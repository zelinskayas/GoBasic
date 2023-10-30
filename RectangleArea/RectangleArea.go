package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("Ожидается ввод длины сторон прямоугольника (a, b): ")
	fmt.Scan(&a, &b)
	fmt.Println("Площадь прямоугольника равна: ", a*b)
}

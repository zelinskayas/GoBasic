package main

import "fmt"

func Calc(a int, b int, action string) int {
	var res int
	switch action {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	}
	return res
}
func main() {
	var a, b int
	var action string

	fmt.Println("Введите два числа и действие:")
	fmt.Scan(&a, &action, &b)
	fmt.Println("Результат: ", Calc(a, b, action))
}

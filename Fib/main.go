package main

import "fmt"

var fibresmap = make(map[int]int)

func CalcFib(num int) int {
	if num == 0 || num == 1 {
		return num
	}

	if value, ok := fibresmap[num]; ok {
		return value
	} else {
		res := CalcFib(num-1) + CalcFib(num-2)
		fibresmap[num] = res
		return res
	}
}

func main() {
	var num int
	fmt.Println("Введите номер числа Фибоначчи:")
	fmt.Scan(&num)
	fmt.Println("Число Фибоначчи равняется: ", CalcFib(num))
}

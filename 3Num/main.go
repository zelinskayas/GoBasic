package main

import (
	"fmt"
)

func main() {
	var num int

	fmt.Println("Введите трехзначное число: ")
	fmt.Scan(&num)
	//или можно проверить длину числа переведенного в строку strconv.Itoa(num)
	if !(num >= 100 && num < 999) {
		panic("ошибка! число должно быть трехзначное")
	}

	fmt.Println("Кол-во сотен, десятков и единиц в этом числе: ", num/100, num%100/10, num%10)

	//с переводом в строку и срез рун
	//for _, value := range []rune(strconv.Itoa(num)) {
	//	fmt.Println(string(value))
	//}
}

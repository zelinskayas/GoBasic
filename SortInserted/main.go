package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите набор целых чисел через пробел:")

	input, _ := reader.ReadString('\n')
	numsStr := strings.Fields(input)
	var nums []int

	for _, value := range numsStr {
		valint, _ := strconv.Atoi(value)
		nums = append(nums, valint)
	}

	/*
		    Сортировка вставками (Insertion Sort) - это алгоритм сортировки,
			в котором элементы входного массива поочередно выбираются и вставляются
			в отсортированную последовательность элементов. Каждый новый элемент
			сравнивается с уже отсортированными элементами, и вставляется в нужное
			место в последовательности. Этот процесс продолжается до тех пор,
			пока все элементы не будут отсортированы.
	*/

	for i := 0; i < len(nums); i++ {
		j := i
		for ; j >= 1 && nums[j] < nums[j-1]; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}

	fmt.Println(nums)
}

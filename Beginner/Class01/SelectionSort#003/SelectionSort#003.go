package main

import "fmt"

func SelectionSort(num []int) {
	if len(num) < 2 {
		return
	}

	len := len(num)
	minValueIndex := 0

	for i := 0; i < len; i++ {
		for j := i; j < len; j++ {
			minValueIndex = i
			if num[j] < num[minValueIndex] {
				minValueIndex = j
			}
		}
		swap(num, i, minValueIndex)
	}
}

func swap(num []int, i int, j int) {
	temp := num[i]
	num[i] = num[j]
	num[j] = temp
}

func PrintSlice(arr []int) {
	for _, v := range arr {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}

func main() {
	arr := []int{3, 17, 8, 39, 100, 15, 12, 2}

	PrintSlice(arr)
	SelectionSort(arr)
	PrintSlice(arr)
}

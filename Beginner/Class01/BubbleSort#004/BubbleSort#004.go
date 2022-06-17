package main

import "fmt"

func PrintSlice(arr []int) {
	for _, v := range arr {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}

func swap(num []int, i int, j int) {
	temp := num[i]
	num[i] = num[j]
	num[j] = temp
}

func BubbleSort(num []int) {
	if len(num) < 2 {
		return
	}

	len := len(num)

	for end := len; end >= 0; end-- {
		for i := 1; i < end; i++ {
			if num[i] < num[i-1] {
				swap(num, i, i-1)
			}
		}
	}
}

func main() {
	arr := []int{3, 17, 8, 39, 100, 15, 12, 2}

	PrintSlice(arr)
	BubbleSort(arr)
	PrintSlice(arr)
}

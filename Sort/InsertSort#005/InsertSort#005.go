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

func InsertSort(num []int) {
	if len(num) < 2 {
		return
	}

	len := len(num)

	for end := 1; end < len; end++ {
		newNumIndex := end
		for ; newNumIndex-1 >= 0 && num[newNumIndex] < num[newNumIndex-1]; newNumIndex-- {
			swap(num, newNumIndex, newNumIndex-1)
		}
	}
}

func main() {
	arr := []int{3, 17, 8, 39, 100, 15, 12, 2}

	PrintSlice(arr)
	InsertSort(arr)
	PrintSlice(arr)
}

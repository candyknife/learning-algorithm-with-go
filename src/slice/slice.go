package slice

import "fmt"

func Print(arr []int) {
	for _, v := range arr {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}

func Swap(num []int, i int, j int) {
	temp := num[i]
	num[i] = num[j]
	num[j] = temp
}

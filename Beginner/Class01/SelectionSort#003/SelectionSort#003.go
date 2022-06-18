/* ————————————————————————————————————————
问题类别：排序
问题描述：给定数组array，对其进行选择排序
时间复杂度：
空间复杂度：
———————————————————————————————————————— */
package main

import slc "learning-algorithm-with-go/src/slice"

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
		slc.Swap(num, i, minValueIndex)
	}
}

func main() {
	arr := []int{3, 17, 8, 39, 100, 15, 12, 2}

	slc.Print(arr)
	SelectionSort(arr)
	slc.Print(arr)
}

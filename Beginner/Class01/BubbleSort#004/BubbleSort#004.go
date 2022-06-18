/* ————————————————————————————————————————
问题类别：排序
问题描述：给定数组array，对其进行冒泡排序
时间复杂度：
空间复杂度：
———————————————————————————————————————— */
package main

import slc "learning-algorithm-with-go/src/slice"

func BubbleSort(num []int) {
	if len(num) < 2 {
		return
	}

	len := len(num)

	for end := len; end >= 0; end-- {
		for i := 1; i < end; i++ {
			if num[i] < num[i-1] {
				slc.Swap(num, i, i-1)
			}
		}
	}
}

func main() {
	arr := []int{3, 17, 8, 39, 100, 15, 12, 2}

	slc.Print(arr)
	BubbleSort(arr)
	slc.Print(arr)
}

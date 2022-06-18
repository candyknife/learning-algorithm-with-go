/* ————————————————————————————————————————
问题类别：排序
问题描述：给定数组array，对其进行插入排序
时间复杂度：
空间复杂度：
———————————————————————————————————————— */
package main

import slc "learning-algorithm-with-go/src/slice"

func InsertSort(num []int) {
	if len(num) < 2 {
		return
	}

	len := len(num)

	for end := 1; end < len; end++ {
		for pre := end; pre-1 >= 0 && num[pre] < num[pre-1]; pre-- {
			slc.Swap(num, pre, pre-1)
		}
	}
}

func main() {
	arr := []int{3, 17, 8, 39, 100, 15, 12, 2}

	slc.Print(arr)
	InsertSort(arr)
	slc.Print(arr)
}

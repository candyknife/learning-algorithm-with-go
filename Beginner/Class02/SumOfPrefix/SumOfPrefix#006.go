/* ————————————————————————————————————————
问题类别：
问题描述：给定数组array，建立其前缀和数组，用于实现快速查询sum(arr, L, R)
时间复杂度：n
空间复杂度：n^2 / 2
———————————————————————————————————————— */
package main

import slc "learning-algorithm-with-go/src/slice"

func SumOfPrefix(arr []int) []int {
	var Help []int
	Help = append(Help, arr[0])
	for i := 1; i < len(arr); i++ {
		Help = append(Help, Help[i-1]+arr[i])
	}
	return Help
}

func main() {

	num := []int{3, 7, 9, 2, 1, 5, 4, 8}

	Help := SumOfPrefix(num)

	slc.Print(Help)

}

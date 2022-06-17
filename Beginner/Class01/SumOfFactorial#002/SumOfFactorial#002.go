/* ————————————————————————————————————————
问题类别：
问题描述：给定数字n，计算 1!+2!+...+n!
时间复杂度：N/A
空间复杂度：N/A
———————————————————————————————————————— */
package main

func SumOfFactorial(num int) int {
	sum, cur := 1, 1
	for i := 2; i <= num; i++ {
		cur = cur * i
		sum += cur
	}
	return sum
}

func main() {

	num := 3

	sum := SumOfFactorial(num)

	println(sum)

}

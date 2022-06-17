/* ————————————————————————————————————————
问题类别：位运算
问题描述：给定 int32 类型的数字，打印其二进制形式
时间复杂度：N/A
空间复杂度：N/A
———————————————————————————————————————— */
package main

import "fmt"

func PrintBinary(num int) {

	for i := 31; i >= 0; i-- {
		if num&(1<<i) == 0 {
			fmt.Print(0)
		} else {
			fmt.Print(1)
		}
	}

	fmt.Println()

}

func main() {

	num := 1024
	PrintBinary(num)
	PrintBinary(num << 1)
	PrintBinary(num >> 1)
	fmt.Printf("\nnum's type is %T", num)

}

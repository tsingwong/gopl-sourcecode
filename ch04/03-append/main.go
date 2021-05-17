/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-18 07:16:43
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-18 07:50:05
 */
package main

import "fmt"

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap = %d \t%v\n", i, cap(y), y)
		x = y
	}
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	// 检测底层数组是否有足够容量来保存新添加的元素
	if zlen <= cap(x) {
		// 有足够空间，拓展 slice，仍然在原来的底层数组上
		z = x[:zlen]
	} else {
		// 没有足够的增长空间（也就会 cap 容量）
		// 优先分配 大于等于 2倍长度的底层数组
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

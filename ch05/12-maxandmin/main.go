/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-06-04 07:11:12
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-06-04 07:48:13
 */
package main

import "fmt"

func main() {
	fmt.Println(min(3, -1, 4))
	fmt.Println(max(3, -1, 4))
}

func min(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	result := nums[0]
	for _, num := range nums {
		if num < result {
			result = num
		}
	}
	return result
}

func max(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	result := nums[0]
	for _, num := range nums {
		if num > result {
			result = num
		}
	}
	return result
}

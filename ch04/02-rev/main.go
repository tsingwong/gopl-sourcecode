/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-17 23:34:59
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-17 23:56:06
 */
package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
func main() {
	a := [5]int{1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a) // [5,4,3,2,1]
	b := make([]int, 5)
	b = append(b, 1, 2, 3, 4, 5)
	rotateLeft(b, 2)
	fmt.Println(b, len(b), cap(b)) // len(b) = cap(b) = 10
	c := []int{1, 2, 3, 4, 5}
	rotateLeft(c, 2)
	fmt.Println(c, len(c), cap(c)) // len(c) = cap(c) = 5
}

// [1,2,3,4,5]
// [2,1,3,4,5]
// [2,1,5,4,3]
// [3,4,5,1,2]
func rotateLeft(s []int, x int) {
	reverse(s[:x])
	reverse(s[x:])
	reverse(s)
}

/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-17 22:09:36
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-17 22:27:39
 */
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	//
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	got := shaBitDiff([]byte{0}, []byte{6})
	fmt.Printf("got %v, want %v", got, 2)
}

func popCount(b byte) int {
	count := 0
	for ; b != 0; count++ {
		b &= b - 1
	}
	return count
}

func bitDiff(a, b []byte) int {
	count := 0
	for i := 0; i < len(a) || i < len(b); i++ {
		switch {
		case i >= len(a):
			count += popCount(b[i])
		case i >= len(b):
			count += popCount(a[i])
		default:
			count += popCount(a[i] ^ b[i])
		}
	}
	return count
}

func shaBitDiff(a, b []byte) int {
	shaA := sha256.Sum256(a)
	shaB := sha256.Sum256(b)
	return bitDiff(shaA[:], shaB[:])
}

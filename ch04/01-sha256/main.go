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
}

/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-06-02 22:37:08
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-06-02 22:44:31
 */
package main

import "fmt"

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

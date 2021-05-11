/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-11 18:00:51
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-11 18:02:22
 */
package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g °F or %g°C\n", f, c)
}

/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-11 19:35:42
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-11 19:39:03
 */
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C", freezingF, fToC(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

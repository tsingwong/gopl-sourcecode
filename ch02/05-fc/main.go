/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-12 22:56:06
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-12 23:32:55
 */
package main

import (
	"fmt"
	tempconv "gopl/ch02/04-tempconv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		// 不能用 %s 可能是因为前面没有给他定义 String 方法
		// fmt.Printf("%v = %v\n %v = %v\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
		fmt.Printf("%s = %s\n %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}

/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-08 11:14:11
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-08 11:53:06
 */
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	fmt.Println("当前命令名称为：", os.Args[0])
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

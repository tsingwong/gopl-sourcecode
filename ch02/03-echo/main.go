/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-11 20:26:06
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-11 20:32:55
 */
package main

import (
	"flag"
	"fmt"
	"strings"
)

/**
 * 创建一个新的对应布尔行标示参数的变量
 * 第一个参数命令行标示参数
 * 第二个参数默认值
 * 第三个参数描述信息
 * 需要注意这里返回都是指针
 */
var n = flag.Bool("n", false, "omit trailing newline")

// 创建一个对应字符串类型的标示参数
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

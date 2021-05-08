/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-08 11:45:33
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-08 11:46:32
 */
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

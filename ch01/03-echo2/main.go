/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-08 11:37:58
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-08 11:57:41
 */
package main

import (
	"fmt"
	"os"
)

func main() {

	for key, value := range os.Args[1:] {
		fmt.Println(key, value)
	}

}

/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-08 12:01:27
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-08 16:33:26
 */
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d \t %s \n", n, line)
		}
	}
	// cmd + D  结束
}

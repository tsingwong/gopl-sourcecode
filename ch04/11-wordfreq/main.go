/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-19 23:36:29
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-19 23:44:05
 */
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	freq := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	// 在第一次调用Scan前先调用input.Split(bufio.ScanWords)函数
	// 这样可以按单词而不是按行输入。
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		freq[word]++
	}
	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, scanner.Err())
		os.Exit(1)
	}
	for word, n := range freq {
		fmt.Printf("%-30s %d\n", word, n)
	}
}

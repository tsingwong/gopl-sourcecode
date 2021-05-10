/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-08 18:30:07
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-09 00:06:58
 */
package main

import (
	"bufio"
	"fmt"
	"os"
)

type obj struct {
	filename string
	count    int
}

func main() {
	counts := make(map[string]*obj)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "none")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				// 标准输出流中打印一条信息
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n.count > 1 {
			fmt.Printf("%d\t %s\t%s\n", n.count, n.filename, line)
		}
	}
}

func countLines(f *os.File, counts map[string]*obj, filename string) {
	input := bufio.NewScanner(f)
	fmt.Println(counts)
	for input.Scan() {
		if _, ok := counts[input.Text()]; !ok {
			counts[input.Text()] = &obj{
				filename: filename,
				count:    1,
			}
		} else {
			counts[input.Text()].count++
		}

	}
}

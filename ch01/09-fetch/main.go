/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-10 13:43:12
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-10 14:10:11
 */
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		httPrefix := "http://"
		if !strings.HasPrefix(url, httPrefix) {
			url = httPrefix + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// 从 response 中读取全部内容输出到 标准输出中
		_, err = io.Copy(os.Stdout, resp.Body)
		fmt.Printf("HTTP status: %d\n", resp.StatusCode)
		// 防止资源泄露，关闭流
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %v\n", err)
			os.Exit(1)
		}
	}
}

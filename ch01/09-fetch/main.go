/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-10 13:43:12
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-10 13:59:26
 */
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// 从 response 中读取全部内容
		b, err := ioutil.ReadAll(resp.Body)
		// 防止资源泄露，关闭流
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", b)
	}
}

/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-11 09:49:29
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-11 10:11:41
 */
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		// 启动 goroutine
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		// 接受 channel 发送来的数据
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		// 通过 channel 发送数据
		ch <- fmt.Sprint(err)
		return
	}
	// ioutil.Discard 可以把他看作一个垃圾桶，可以向里面写一些不需要的数据
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("While reading %s : %v", url, err)
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2f %7d %s", secs, nbytes, url)
}

/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-10 23:38:26
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-10 23:42:36
 */
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.path = %q\n", r.URL.Path)
}

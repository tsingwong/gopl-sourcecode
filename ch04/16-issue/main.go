/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-23 19:12:40
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-24 22:25:57
 */
package main

import (
	"fmt"
	github "gopl/ch04/15-github"
	"log"
	"os"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-23 19:12:40
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-24 22:46:30
 */
package main

import (
	"fmt"
	github "gopl/ch04/15-github"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	// for _, item := range result.Items {
	// 	fmt.Printf("#%-5d %9.9s %.55s\n",
	// 		item.Number, item.User.Login, item.Title)
	// }

	format := "#%-5d %9.9s %.55s\n"

	now := time.Now()
	pastDay := make([]*github.Issue, 0)
	pastMonth := make([]*github.Issue, 0)
	pastYear := make([]*github.Issue, 0)

	day := now.AddDate(0, 0, -1)
	month := now.AddDate(0, -1, 0)
	year := now.AddDate(-1, 0, 0)

	for _, item := range result.Items {
		switch {
		case item.CreatedAt.After(day):
			pastDay = append(pastDay, item)
		case item.CreatedAt.After(month) && item.CreatedAt.Before(day):
			pastMonth = append(pastMonth, item)
		case item.CreatedAt.After(year) && item.CreatedAt.Before(month):
			pastYear = append(pastYear, item)
		}
	}

	if len(pastDay) > 0 {
		fmt.Printf("\nPast day:\n")
		for _, item := range pastDay {
			fmt.Printf(format, item.Number, item.User.Login, item.Title)
		}
	}
	if len(pastMonth) > 0 {
		fmt.Printf("\nPast month:\n")
		for _, item := range pastMonth {
			fmt.Printf(format, item.Number, item.User.Login, item.Title)
		}
	}
	if len(pastYear) > 0 {
		fmt.Printf("\nPast year:\n")
		for _, item := range pastYear {
			fmt.Printf(format, item.Number, item.User.Login, item.Title)
		}
	}
}

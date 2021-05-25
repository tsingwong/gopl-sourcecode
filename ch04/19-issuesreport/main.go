/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-25 22:39:59
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-25 23:06:04
 */
package main

import (
	github "gopl/ch04/15-github"
	"html/template"
	"log"
	"os"
	"time"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

// 模版通常是在编译时就测试好了，如果模版解析失败那么是致命的。
// template.Must 辅助函数简化知名错误的处理
// 接受一个模版和 error 类型的参数
// 判断 error 是否为 nil，如果不是 nil 就发出 panic 异常
// 然后返回传入的模版
var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func main() {

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	// os.Stdout 作为输出源来执行模版
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-26 16:12:16
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-26 16:18:23
 */
package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	const templ = `<p>a:{{.A}}</p><p>b:{{.B}}</p>`
	t := template.Must(template.New("escape").Parse(templ))
	var data struct {
		A string
		B template.HTML
	}
	data.A = "<b>Hello!</b>"
	data.B = "<b>Hello!</b>"
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}

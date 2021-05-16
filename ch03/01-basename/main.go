/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-14 20:03:51
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-16 16:53:55
 */
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(basename("a/b/c.go")) // "c"
	fmt.Println(basename("c.d.go"))   // "c.d"
	fmt.Println(basename("abc"))      // "abc"

	fmt.Println(basename2("a/b/c.go")) // "c"
	fmt.Println(basename2("c.d.go"))   // "c.d"
	fmt.Println(basename2("abc"))      // "abc"

	fmt.Println(comma("1234567890"))
	fmt.Println(comma("123"))

	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"

	for i := 1; i < len(os.Args[1:]); i++ {
		fmt.Printf("%s\n", comma(os.Args[i]))
	}

}

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

// comma 将表示整数的字符串，每隔三个字符插入要给逗号分隔符
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func unRecursionComma(s string) string {
	var buf bytes.Buffer
	pre := len(s) % 3
	if pre == 0 {
		pre = 3
	}
	buf.WriteString(s[:pre])
	for i := pre; i < len(s); i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}

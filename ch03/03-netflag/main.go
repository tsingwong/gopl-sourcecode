/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-16 22:14:24
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-16 22:45:14
 */
package main

import "fmt"

type Flags uint

const (
	FlagUp           Flags = 1 << iota // 0000 0001
	FlagBroadcast                      // 0000 0010
	FlagLoopback                       //  0000 0100
	FlagPointToPoint                   // 0000 1000
	FlagMulticast                      // 0001 0000
)

func main() {
	var v Flags = FlagMulticast | FlagUp // 0001 0001
	fmt.Printf("%b %t\n", v, IsUp(v))    // "10001 true"
	TurnDown(&v)                         // 0001 0000
	fmt.Printf("%b %t\n", v, IsUp(v))    // "10000 false"
	SetBroadcast(&v)                     // 0001 0010
	fmt.Printf("%b %t\n", v, IsUp(v))    // "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v))  // "10010 true"
}

func IsUp(v Flags) bool {
	return v&FlagUp == FlagUp
}

func IsCast(v Flags) bool {
	return v&(FlagBroadcast|FlagMulticast) != 0
}

func TurnDown(v *Flags) {
	// 按位置零，左边与右边相同的位置清零，不同的位置保留
	*v &^= FlagUp
}

func SetBroadcast(v *Flags) {
	*v |= FlagBroadcast
}

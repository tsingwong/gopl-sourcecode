/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-13 13:26:20
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-13 13:30:10
 */
package popcount

import (
	"testing"
)

func bench(b *testing.B, f func(uint64) int) {
	for i := 0; i < b.N; i++ {
		f(uint64(i))
	}
}

// BenchmarkTable-8   	1000 000 000	         0.2875 ns/op	       0 B/op	       0 allocs/op
func BenchmarkTable(b *testing.B) {
	bench(b, PopCount)
}

// BenchmarkTableLoop-8   	382 160 126	         3.133 ns/op	       0 B/op	       0 allocs/op
func BenchmarkTableLoop(b *testing.B) {
	bench(b, PopCountLoop)
}

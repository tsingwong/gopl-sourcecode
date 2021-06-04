/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-06-03 16:32:31
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-06-04 11:41:24
 */
package sum

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		got  []int
		want int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{1, 2, 3, 4}, 10},
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{1}, 1},
	}
	for _, test := range tests {
		if sum(test.got...) != test.want {
			t.Errorf("Sum(%v): got: %v, want: %v ", test.got, sum(test.got...), test.want)
		}
	}
}

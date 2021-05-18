/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-18 23:21:21
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-19 07:08:28
 */
package reverse

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		a, want [5]int
	}{
		{[5]int{1, 2, 3, 4, 5}, [5]int{5, 4, 3, 2, 1}},
	}
	for _, test := range tests {
		reverse(&test.a)

		if !reflect.DeepEqual(test.a, test.want) {
			t.Errorf("got %v, want %v", test.a, test.want)
		}
	}
}

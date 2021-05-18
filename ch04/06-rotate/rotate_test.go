/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-19 07:13:32
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-19 07:18:43
 */
package rotate

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	s := []int{1, 2, 3, 4}
	rotate(s)
	want := []int{2, 3, 4, 1}
	if !reflect.DeepEqual(s, want) {
		t.Errorf("got %v, want %v", s, want)
	}
}

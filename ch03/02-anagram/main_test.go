/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-16 17:27:58
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-16 17:29:27
 */
package anagram

import "testing"

func TestAnagram(t *testing.T) {
	tests := []struct {
		a, b string
		want bool
	}{
		{"aba", "baa", true},
		{"aaa", "baa", false}, // same characters but different frequencies
	}
	for _, test := range tests {
		got := isAnagram(test.a, test.b)
		if got != test.want {
			t.Errorf("isAnagram(%q, %q), got %v, want %v",
				test.a, test.b, got, test.want)
		}
	}
}

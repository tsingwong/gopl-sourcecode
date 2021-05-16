/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-16 17:23:02
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-16 17:28:24
 */
package main

func isAnagram(s1, s2 string) bool {
	s1Freq := make(map[rune]int)
	for _, c := range s1 {
		s1Freq[c]++
	}

	s2Freq := make(map[rune]int)
	for _, c := range s2 {
		s2Freq[c]++
	}

	for k, v := range s1Freq {
		if s2Freq[k] != v {
			return false
		}
	}
	for k, v := range s2Freq {
		if s1Freq[k] != v {
			return false
		}
	}
	return true
}

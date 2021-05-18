/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-18 23:14:24
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-18 23:21:04
 */
package reverse

func reverse(ints *[5]int) {
	for i, j := 0, len(ints)-1; i < j; i, j = i+1, j-1 {
		ints[i], ints[j] = ints[j], ints[i]
	}
}

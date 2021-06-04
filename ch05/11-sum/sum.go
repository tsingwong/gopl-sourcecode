/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-06-03 16:28:46
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-06-03 16:29:00
 */
package sum

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

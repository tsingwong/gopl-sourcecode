/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-19 07:08:25
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-19 07:13:23
 */
package rotate

func rotate(slice []int) {
	first := slice[0]
	copy(slice, slice[1:])
	slice[len(slice)-1] = first
}

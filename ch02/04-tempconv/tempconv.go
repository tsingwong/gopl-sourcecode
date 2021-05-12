/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-12 22:25:08
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-12 22:39:34
 */
package tempconv

type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

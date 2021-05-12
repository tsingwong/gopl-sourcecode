/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-12 22:38:47
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-12 22:38:50
 */
package tempconv

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

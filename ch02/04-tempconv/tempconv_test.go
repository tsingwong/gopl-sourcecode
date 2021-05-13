/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-13 11:17:09
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-13 11:33:39
 */
package tempconv

import (
	"math"
	"testing"
)

func TestTempConv(t *testing.T) {
	tests := []struct {
		f Fahrenheit
		c Celsius
		k Kelvin
	}{
		{68, 20, 293.15},
		{32, 0, 273.15},
		{-40, -40, 233.15},
	}
	eps := 0.0000001 // 误差
	for _, test := range tests {
		if math.Abs(float64(CToF(test.c)-test.f)) > eps {
			t.Errorf("CToF(%s): got %s, want %s", test.c, CToF(test.c), test.f)
		}
		if math.Abs(float64(FToC(test.f)-test.c)) > eps {
			t.Errorf("FToC(%s): got %s, want %s", test.f, FToC(test.f), test.c)
		}
		if math.Abs(float64(KToC(test.k)-test.c)) > eps {
			t.Errorf("KToC(%s): got %s, want %s", test.k, KToC(test.k), test.c)
		}
	}
}

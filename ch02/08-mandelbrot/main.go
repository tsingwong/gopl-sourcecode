/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-13 23:23:55
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-13 23:47:15
 */
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"os"
	"strconv"
)

func main() {
	ops := map[string]float64{
		"width":  1024,
		"height": 1024,
	}
	if len(os.Args) > 1 && os.Args[1] == "web" {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if err := r.ParseForm(); err != nil {
				log.Print(err)
			}
			fmt.Println("改变前", ops)
			for k, v := range r.Form {

				temp, err := strconv.ParseFloat(v[0], 64)
				if err != nil {

					log.Fatalf("err %s", err)
				}
				fmt.Println(k, temp)

				ops[k] = temp
			}
			fmt.Println("改变后", ops)
			genPng(w, ops)
		})
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	genPng(os.Stdout, ops)

}

func genPng(out io.Writer, ops map[string]float64) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
	)
	var width, height float64 = 1024, 1024
	if value, ok := ops["width"]; ok {
		width = value
	}
	if value, ok := ops["height"]; ok {
		height = value
	}

	img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	for py := 0; py < int(height); py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < int(width); px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(out, img)
}

func mandelbrot(z complex128) color.Color {
	const (
		iterations = 200
		contrast   = 15
	)

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

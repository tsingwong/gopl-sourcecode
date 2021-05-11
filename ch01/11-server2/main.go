/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-05-10 23:51:29
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-05-11 08:38:12
 */
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	ops := map[string]float64{
		"cycles":  5,
		"res":     0.001,
		"size":    100,
		"nframes": 64,
		"delay":   8,
	}
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/lissajous", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		for k, v := range r.Form {
			temp, err := strconv.Atoi(v[0])
			if err != nil {
				ops[k] = float64(temp)
			}
		}
		lissajous(w, ops)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func lissajous(out io.Writer, ops map[string]float64) {

	cycles := int(ops["cycles"])
	res := ops["res"]
	size := int(ops["size"])
	nframes := int(ops["nframes"])
	delay := int(ops["delay"])

	palette := make([]color.Color, 0, nframes)
	palette = append(palette, color.RGBA{0, 0, 0, 255})

	for i := 1; i < nframes; i++ {
		scale := float64(i) / float64(nframes)
		c := color.RGBA{uint8(55 + 200*scale), uint8(55 + 200*scale), uint8(55 + 200*scale), 255}
		palette = append(palette, c)
	}

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
				uint8(i%len(palette)-1)+1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

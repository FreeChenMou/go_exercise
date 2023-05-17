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
)

var palette = []color.Color{color.RGBA{0x00, 0x1a, 0x2b, 0xff},
	color.RGBA{0x1e, 0x2d, 0xcd, 0xff},
	color.RGBA{0x4e, 0x5d, 0x7d, 0xff},
	color.RGBA{0x00, 0x2d, 0x7d, 0xff},
}

const (
	//cycles  = 5     // number of complete x oscillator revolutions
	res     = 0.001 // angular resolution
	size    = 100   // image canvas covers [-size..+size]
	nframes = 64    // number of animation frames
	delay   = 8     // delay between frames in 10ms units
)

//exercise1.12 获取请求Url参数
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	cycles := 5

	r.ParseForm()
	if len(r.Form["cycle"]) > 0 {
		s := r.Form["cycle"][0]
		atom, err := strconv.Atoi(s)
		if err != nil {
			fmt.Fprintf(w, "server:%v \n", err)
		}
		if atom > 0 {
			cycles = atom
		}
	}

	lissajous(w, cycles)
}

func lissajous(out io.Writer, cycles int) {

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(i%4))
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

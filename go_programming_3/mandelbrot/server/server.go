package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

var palette = []color.Color{
	color.RGBA{0xab, 0xbb, 0xca, 0xff},
}

//exercise3.9
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	const (
		width, height = 1024, 1024
	)
	var xmin, ymin, xmax, ymax float64 = -2, -2, +2, +2
	if len(request.Form["x"]) > 0 {
		s := request.Form["x"][0]
		x, err := strconv.Atoi(s)
		if err != nil {
			fmt.Fprintf(writer, "server:%v \n", err)
		}
		xmax, xmin = float64(+x), float64(-x)
	}

	if len(request.Form["y"]) > 0 {
		s := request.Form["y"][0]
		y, err := strconv.Atoi(s)
		if err != nil {
			fmt.Fprintf(writer, "server:%v \n", err)
		}
		ymax, ymin = float64(+y), float64(-y)
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			//SubPixel := make([]color.Color, 0)
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)

			img.Set(px, py, mandelbrot(z))
		}
		// Image point (px, py) represents complex value z.

	}
	png.Encode(writer, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return palette[0]
		}
	}
	return color.Black
}

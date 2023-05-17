package main

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"math/big"
	"math/cmplx"
)

var palette = [16]color.Color{
	color.RGBA{66, 30, 15, 255},    // # brown 3
	color.RGBA{25, 7, 26, 255},     // # dark violett
	color.RGBA{9, 1, 47, 255},      //# darkest blue
	color.RGBA{4, 4, 73, 255},      //# blue 5
	color.RGBA{0, 7, 100, 255},     //# blue 4
	color.RGBA{12, 44, 138, 255},   //# blue 3
	color.RGBA{24, 82, 177, 255},   //# blue 2
	color.RGBA{57, 125, 209, 255},  //# blue 1
	color.RGBA{134, 181, 229, 255}, // # blue 0
	color.RGBA{211, 236, 248, 255}, // # lightest blue
	color.RGBA{241, 233, 191, 255}, // # lightest yellow
	color.RGBA{248, 201, 95, 255},  // # light yellow
	color.RGBA{255, 170, 0, 255},   // # dirty yellow
	color.RGBA{204, 128, 0, 255},   // # brown 0
	color.RGBA{153, 87, 0, 255},    // # brown 1
	color.RGBA{106, 52, 3, 255},    // # brown 2
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		pixelX, pixelY         = (xmax - xmin) / width, (ymax - ymin) / height
	)

	//cmpsateX := []float64{-pixelX, pixelX} exercise 3.6
	//cmpsateY := []float64{-pixelY, pixelY}
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			//SubPixel := make([]color.Color, 0)
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)

			img.Set(px, py, mandelbrotRat(z))
		}
		// Image point (px, py) represents complex value z.

	}
	buf := &bytes.Buffer{}
	png.Encode(buf, img) // NOTE: ignoring errors
	if err := ioutil.WriteFile("out.png", buf.Bytes(), 0666); err != nil {
		panic(err)
	}
}

func avg(pixel []color.Color) color.Color {
	var r, g, b, a uint16
	length := len(pixel)

	for _, c := range pixel {
		r1, g1, b1, a1 := c.RGBA()
		r += uint16(r1 / uint32(length))
		g += uint16(g1 / uint32(length))
		b += uint16(b1 / uint32(length))
		a += uint16(a1 / uint32(length))
	}
	return color.RGBA64{r, g, b, a}
}

//exercise3.5
func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//exercise3.7 牛顿法
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}

//exercise3.8 精度有限和无限精度的使用
func mandelbrot64(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + complex64(z)
		if cmplx.Abs(complex128(v)) > 2 {
			return getColor(n)
		}
	}
	return color.Black
}

func mandelbrotBigFloat(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	zR := (&big.Float{}).SetFloat64(real(z))
	zI := (&big.Float{}).SetFloat64(imag(z))
	var vR, vI = &big.Float{}, &big.Float{}
	for i := uint8(0); i < iterations; i++ {
		vR2, vI2 := &big.Float{}, &big.Float{}
		vR2.Mul(vR, vR).Sub(vR2, (&big.Float{}).Mul(vI, vI)).Add(vR2, zR)
		vI2.Mul(vR, vI).Mul(vI2, big.NewFloat(2)).Add(vI2, zI)
		vR, vI = vR2, vI2
		squareSum := &big.Float{}
		squareSum.Mul(vR, vR).Add(squareSum, (&big.Float{}).Mul(vI, vI))
		if squareSum.Cmp(big.NewFloat(4)) == 1 {
			return getColor(i)
		}
	}
	return color.Black
}

func mandelbrotRat(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	zR := (&big.Rat{}).SetFloat64(real(z))
	zI := (&big.Rat{}).SetFloat64(imag(z))
	var vR, vI = &big.Rat{}, &big.Rat{}
	for i := uint8(0); i < iterations; i++ {
		// (r+i)^2 = r^2 + 2ri + i^2
		vR2, vI2 := &big.Rat{}, &big.Rat{}
		vR2.Mul(vR, vR).Sub(vR2, (&big.Rat{}).Mul(vI, vI)).Add(vR2, zR)
		vI2.Mul(vR, vI).Mul(vI2, big.NewRat(2, 1)).Add(vI2, zI)
		vR, vI = vR2, vI2
		squareSum := &big.Rat{}
		squareSum.Mul(vR, vR).Add(squareSum, (&big.Rat{}).Mul(vI, vI))
		if squareSum.Cmp(big.NewRat(4, 1)) == 1 {
			return getColor(i)
		}
	}
	return color.Black
}

func getColor(i uint8) color.Color {
	return palette[i%16]
}

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strings"
)

const (
	width, height = 800, 400            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y
)

var sin, cos = math.Sin(angle), math.Cos(angle)

func main() {
	http.HandleFunc("/", lissajous)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

//exercise 3.4
func lissajous(w http.ResponseWriter, request *http.Request) {
	var result string
	var builder strings.Builder
	w.Header().Set("Content-Type", "image/svg+xml")
	builder.WriteString(fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height))

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)

			//exercise 3.3
			if az < 0 && bz < 0 && cz < 0 && dz < 0 {
				builder.WriteString(fmt.Sprintf("<polygon style='fill: blue' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy))
			} else {
				builder.WriteString(fmt.Sprintf("<polygon style='fill: red' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy))
			}
		}
	}
	builder.WriteString(fmt.Sprintf("</svg>"))
	result = builder.String()
	_, err := w.Write([]byte(result))
	if err != nil {
		fmt.Fprintf(os.Stderr, "surface err:%v ", err)
		return
	}
}

//exercise5.6 裸返回
func corner(i, j int) (sx float64, sy float64, z float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z = f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx = width/2 + (x-y)*cos*xyscale
	sy = height/2 + (x+y)*sin*xyscale - z*zscale
	return sx, sy, z
}

//exercise3.1 处理无穷大或Nan值
func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	//exercise 3.2
	ans := math.Atanh(r) / r
	if math.IsNaN(ans) || math.IsInf(ans, 0) {
		return math.Sin(r)
	}
	return ans
}

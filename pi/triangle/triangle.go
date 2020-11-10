package main

import (
	"flag"
	"fmt"
	"math"
	"time"

	c "github.com/lunatikub/lunamath/common"
)

var radius = 1.0

// emulate a manual measure with a guide of the triangle base
func measureBase(a, theta float64) float64 {
	return 2 * a * math.Cos(c.Radian(theta))
}

// emulate a manual measure with a guide of the triangle height
func measureHeight(a, theta float64) float64 {
	return a * math.Sin(c.Radian(theta))
}

func trianglePIEstimation(triangles float64) float64 {
	alpha := 360.0 / triangles
	theta := (180.0 - alpha) / 2.0
	base := measureBase(radius, theta)
	height := measureHeight(radius, theta)
	area := (base * height) / 2.0
	return area * triangles
}

type options struct {
	triangles int
}

type point struct {
	x, y float64
}

func getOptions() *options {
	opts := new(options)
	flag.IntVar(&opts.triangles, "triangle", 18, "number of triangles")
	flag.Parse()
	return opts
}

const (
	sz = 1000
	xo = sz / 2 // x ogirin
	yo = sz / 2 // y origin
)

func convert(x, y float64) (int, int) {
	return int(x*float64(sz/2.0)) + xo, int(y*float64(sz/2)) + yo
}

func main() {
	opts := getOptions()
	S := c.SDLInit(sz, sz)

	delta := 360.0 / float64(opts.triangles)
	angle := 0.0

	x := 0
	y := 0
	prevX := 0
	prevY := 0
	first := true

	S.Circle(xo, yo, sz/2, c.Red)
	for {
		x, y = convert(math.Sin(c.Radian(angle)),
			math.Cos(c.Radian(angle)))
		S.Line(xo, yo, x, y, c.Green)
		angle += delta
		if !first {
			S.Line(prevX, prevY, x, y, c.Blue)
		}
		prevX, prevY = x, y
		if angle > 360 {
			break
		}
		S.Refresh()
		wait := time.Duration(1000 / opts.triangles)
		time.Sleep(time.Millisecond * wait)
		first = false
	}
	S.Line(prevX, prevY, x, y, c.Blue)
	S.Refresh()
	fmt.Println(trianglePIEstimation(float64(opts.triangles)))
	S.Wait()
}

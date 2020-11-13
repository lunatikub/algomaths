package main

import (
	"flag"
	"fmt"
	"math"
	"time"

	common "github.com/lunatikub/lunamath/common"
)

const (
	sz     = 1000
	xo     = sz / 2 // x ogirin
	yo     = sz / 2 // y origin
	radius = 1.0
)

// emulate a manual measure with a guide of the triangle base
func measureBase(a, theta float64) float64 {
	return 2 * a * math.Cos(common.Radian(theta))
}

// emulate a manual measure with a guide of the triangle height
func measureHeight(a, theta float64) float64 {
	return a * math.Sin(common.Radian(theta))
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

func getOptions() *options {
	opts := new(options)
	flag.IntVar(&opts.triangles, "triangle", 18, "number of triangles")
	flag.Parse()
	if opts.triangles < 3 {
		panic("minimal number of triangles required: 3")
	}
	return opts
}

func convert(x, y float64) (int, int) {
	return int(x*float64(sz/2.0)) + xo, int(y*float64(sz/2)) + yo
}

func main() {
	opts := getOptions()
	S := common.SDLInit(sz, sz)

	delta := 360.0 / float64(opts.triangles)
	angle := 0.0

	x := 0
	y := 0
	prevX := 0
	prevY := 0
	once := true

	S.Circle(xo, yo, sz/2, common.Red)
	for {
		x, y = convert(math.Sin(common.Radian(angle)),
			math.Cos(common.Radian(angle)))
		S.Line(xo, yo, x, y, common.Green)
		angle += delta
		if !once {
			S.Line(prevX, prevY, x, y, common.Blue)
		}
		prevX, prevY = x, y
		if angle > 360.0 {
			break
		}
		S.Refresh()
		time.Sleep(time.Millisecond *
			time.Duration(1000/opts.triangles))
		once = false
	}
	S.Line(prevX, prevY, x, y, common.Blue)
	S.Refresh()
	fmt.Println(trianglePIEstimation(float64(opts.triangles)))
	S.Wait()
}

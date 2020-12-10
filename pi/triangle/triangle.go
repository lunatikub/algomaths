package main

import (
	"flag"
	"fmt"
	"math"
	"time"

	common "github.com/lunatikub/AlgoMathAndCo/common"
)

const (
	sz     = 1000
	margin = 10
	xo     = sz / 2 // x ogirin
	yo     = sz / 2 // y origin
	radius = 1.0
)

// emulate a manual measure with a ruler of the triangle base
func measureBase(a, theta float64) float64 {
	return 2 * a * math.Cos(common.Radian(theta))
}

func trianglePIEstimation(triangles float64) float64 {
	alpha := 360.0 / triangles
	theta := (180.0 - alpha) / 2.0
	base := measureBase(radius, theta)
	circumference := base * triangles
	return circumference / (2 * radius)
}

type options struct {
	triangles int
	thickness int
}

func getOptions() *options {
	opts := new(options)
	flag.IntVar(&opts.triangles, "triangle", 18, "number of triangles")
	flag.IntVar(&opts.thickness, "thickness", 2, "thickness")
	flag.Parse()
	if opts.triangles < 3 {
		panic("minimal number of triangles required: 3")
	}
	return opts
}

func convert(x, y float64) (int, int) {
	return int(x*float64(sz/2.0)) + xo + margin,
		int(y*float64(sz/2)) + yo + margin
}

func main() {
	opts := getOptions()
	S := common.SDLInit(sz+2*margin, sz+2*margin)

	delta := 360.0 / float64(opts.triangles)
	angle := 0.0

	x := 0
	y := 0
	prevX := 0
	prevY := 0
	once := true

	S.Circle(xo+margin, yo+margin, sz/2, common.Red, opts.thickness)
	for {
		x, y = convert(math.Sin(common.Radian(angle)),
			math.Cos(common.Radian(angle)))
		S.Line(xo, yo, x, y, common.Green, opts.thickness)
		angle += delta
		if !once {
			S.Line(prevX, prevY, x, y, common.Blue, opts.thickness)
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
	S.Line(prevX, prevY, x, y, common.Blue, opts.thickness)
	S.Refresh()
	fmt.Println(trianglePIEstimation(float64(opts.triangles)))
	S.Wait()
}

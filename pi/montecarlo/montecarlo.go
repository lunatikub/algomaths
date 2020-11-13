package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/lunatikub/lunamath/common"
)

const (
	sz     = 1000
	margin = 10
)

type options struct {
	iteration int
	animated  bool
}

func getOptions() *options {
	opts := new(options)
	flag.IntVar(&opts.iteration, "iteration", 1000, "number of iterations")
	flag.BoolVar(&opts.animated, "animated", false, "Animated Monte Carlo")
	flag.Parse()
	return opts
}

func convert(x, y float64) (int, int) {
	return int(x*sz) + margin, sz - int(y*sz) + margin
}

func monteCarlo(iteration int) float64 {
	r := 0
	n := 0
	for {
		x := rand.Float64()
		y := rand.Float64()
		if x*x+y*y <= 1 {
			r++
		}
		n++
		if n == iteration {
			break
		}

	}
	return 4 * float64(r) / float64(n)
}

func decoration(S *common.SDL) {
	x1, y1 := convert(0, 0)
	x2, y2 := convert(0, 1)
	x3, y3 := convert(1, 0)
	x4, y4 := convert(1, 1)
	S.Line(x1, y1, x2, y2, common.Green, 2)
	S.Line(x1, y1, x3, y3, common.Green, 2)
	S.Line(x2, y2, x4, y4, common.Red, 2)
	S.Line(x3, y3, x4, y4, common.Red, 2)
	S.Sector(x1, y1, sz, common.Green, 2)
}

func monteCarloAnimated(S *common.SDL, iteration int) float64 {
	r := 0
	n := 0
	for {
		x, y := rand.Float64(), rand.Float64()
		xp, yp := convert(x, y)
		if x*x+y*y <= 1 {
			S.Set(xp, yp, common.Green, 2)
			r++
		} else {
			S.Set(xp, yp, common.Red, 2)
		}
		S.Refresh()
		n++
		if n == iteration {
			break
		}
	}
	return 4 * float64(r) / float64(n)
}

func main() {
	opts := getOptions()
	rand.Seed(time.Now().UnixNano())
	pi := 0.0
	var S *common.SDL

	if opts.animated {
		S = common.SDLInit(sz+2*margin, sz+2*margin)
		decoration(S)
		pi = monteCarloAnimated(S, opts.iteration)
	} else {
		pi = monteCarlo(opts.iteration)
	}
	fmt.Printf("pi = %v\n", pi)
	if opts.animated {
		S.Wait()
	}
}

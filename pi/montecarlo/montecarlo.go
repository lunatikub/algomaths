package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"time"

	gc "github.com/gbin/goncurses"
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

func winUpdate(win *gc.Window, r, n, e int) {
	pi := 4 * float64(r) / float64(n)
	win.MovePrintf(2, 2, "PI estimation:    %-100v", pi)
	win.MovePrintf(3, 2, "Precision:        %-100v", math.Abs(math.Pi-pi))
	win.MovePrintf(4, 2, "IN (green):       %-100v", r)
	win.MovePrintf(5, 2, "Iterations:       %-100v", n)
	win.MovePrintf(6, 2, "End:              %-100v", e)
	win.Refresh()
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
	S.Line(x1, y1, x2, y2, common.Green)
	S.Line(x1, y1, x3, y3, common.Green)
	S.Line(x2, y2, x4, y4, common.Red)
	S.Line(x3, y3, x4, y4, common.Red)
	S.Sector(x1, y1, sz, common.Green)
}

func monteCarloAnimated(S *common.SDL, W *gc.Window, iteration int) float64 {
	r := 0
	n := 0
	for {
		x, y := rand.Float64(), rand.Float64()
		xp, yp := convert(x, y)
		if x*x+y*y <= 1 {
			S.SetBigPoint(xp, yp, common.Green)
			r++
		} else {
			S.SetBigPoint(xp, yp, common.Red)
		}
		S.Refresh()
		winUpdate(W, r, n, iteration)
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

	if opts.animated {
		nWin := common.NCInit()
		S := common.SDLInit(sz+2*margin, sz+2*margin)
		decoration(S)
		monteCarloAnimated(S, nWin, opts.iteration)
		S.Wait()
	} else {
		fmt.Printf("%v\n", monteCarlo(opts.iteration))
	}
}

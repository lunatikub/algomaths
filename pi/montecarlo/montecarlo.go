package main

import (
	"flag"
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	gc "github.com/gbin/goncurses"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	sz     = 1000
	margin = 10
)

var green = color.RGBA{0, 255, 0, 255}
var red = color.RGBA{255, 0, 0, 255}

type options struct {
	iteration int
	animated  bool
}

type point struct {
	x, y float64
}

func getOptions() *options {
	opts := new(options)
	flag.IntVar(&opts.iteration, "iteration", 1000, "number of iterations")
	flag.BoolVar(&opts.animated, "animated", false, "Animated Monte Carlo")
	flag.Parse()
	return opts
}

func point2SDL(p point) (x, y int) {
	x = int(p.x * sz)
	y = int(p.y * sz)
	x = x + margin
	y = sz - y + margin
	return x, y
}

func line(surface *sdl.Surface, p1, p2 point, C color.RGBA) {
	x1, y1 := point2SDL(p1)
	x2, y2 := point2SDL(p2)

	x := float64(x2 - x1)
	y := float64(y2 - y1)
	len := math.Sqrt(float64(x*x + y*y))
	addx := x / len
	addy := y / len

	x = float64(x1)
	y = float64(y1)
	for i := 0; i < int(len); i++ {
		surface.Set(int(x), int(y), C)
		x += addx
		y += addy
	}
}

func setPoint(surface *sdl.Surface, x, y int, C color.RGBA) {
	var neighbors = []point{
		{0, 1}, {1, 1}, {1, 0}, {1, -1},
		{0, -1}, {-1, -1}, {-1, 0}, {-1, 1},
	}
	surface.Set(x, y, C)
	for _, n := range neighbors {
		surface.Set(x+int(n.x), y+int(n.y), C)
	}
}

func sdlInit() *sdl.Window {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	window, err := sdl.CreateWindow("montecarlo", 20, 20,
		sz+2*margin, sz+2*margin, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

	line(surface, point{0, 0}, point{0, 1}, green)
	line(surface, point{0, 0}, point{1, 0}, green)
	line(surface, point{0, 1}, point{1, 1}, red)
	line(surface, point{1, 1}, point{1, 0}, red)
	window.UpdateSurface()

	return window
}

func sdlExit() {
	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
	}
	sdl.Quit()
}

func ncursesInit() *gc.Window {
	win, err := gc.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer gc.End()

	gc.Echo(false)
	gc.CBreak(true)
	gc.Cursor(0)
	win.Clear()

	return win
}

func screenUpdate(win *gc.Window, r, n, e int) {
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

func monteCarloAnimated(
	sdlWin *sdl.Window,
	nWin *gc.Window,
	iteration int) float64 {
	surface, err := sdlWin.GetSurface()
	if err != nil {
		panic(err)
	}
	r := 0
	n := 0
	for {
		x := rand.Float64()
		y := rand.Float64()
		xSDL, ySDL := point2SDL(point{x, y})
		if x*x+y*y <= 1 {
			setPoint(surface, xSDL, ySDL, green)
			r++
		} else {
			setPoint(surface, xSDL, ySDL, red)
		}
		sdlWin.UpdateSurface()
		screenUpdate(nWin, r, n, iteration)
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
		nWin := ncursesInit()
		sdlWin := sdlInit()
		defer sdlWin.Destroy()
		monteCarloAnimated(sdlWin, nWin, opts.iteration)
		sdlExit()
	} else {
		fmt.Printf("%v\n", monteCarlo(opts.iteration))
	}
}

package common

import (
	"image/color"
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

// SDL structure
type SDL struct {
	win  *sdl.Window
	sur  *sdl.Surface
	w, h int32
}

// Green color
var Green = color.RGBA{0, 255, 0, 255}

// Red color
var Red = color.RGBA{255, 0, 0, 255}

// Blue color
var Blue = color.RGBA{0, 0, 255, 255}

// Black color
var Black = color.RGBA{0, 0, 0, 255}

// White color
var White = color.RGBA{255, 255, 255, 255}

// SDLInit Initialize a SDL window
func SDLInit(w, h int32) *SDL {
	S := SDL{}
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	window, err := sdl.CreateWindow("montecarlo", 20, 20,
		w, h, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)
	window.UpdateSurface()
	S.win = window
	S.sur = surface
	S.w = w
	S.h = h
	return &S
}

// Wait end of SDL window with crtl-c
func (S *SDL) Wait() {
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
	S.win.Destroy()
	sdl.Quit()
}

// Clean the surface
func (S *SDL) Clean() {
	rect := sdl.Rect{X: 0, Y: 0, W: S.w, H: S.h}
	S.sur.FillRect(&rect, 0xffffffff)
}

// Refresh update surface
func (S *SDL) Refresh() {
	S.win.UpdateSurface()
}

type neighbor struct {
	x, y int
}

func (S *SDL) check(x, y int) bool {
	return x >= 0 && y >= 0 && x <= int(S.w) && y <= int(S.h)
}

// Set set a point at coordiantes (x,y)
// with color C and thickness T
func (S *SDL) Set(x, y int, C color.RGBA, T int) {
	switch T {
	case 2:
		var neighbors = []neighbor{
			{0, 1}, {1, 1}, {1, 0}, {1, -1},
			{0, -1}, {-1, -1}, {-1, 0}, {-1, 1},
		}
		if S.check(x, y) {
			S.sur.Set(x, y, C)
		}
		for _, n := range neighbors {
			if S.check(x+n.x, y+n.y) {
				S.sur.Set(x+n.x, y+n.y, C)
			}
		}
	case 1:
		if S.check(x, y) {
			S.sur.Set(x, y, C)
		}
	}
}

// Line Draw a line from (x1, y1) to (x2, y2)
// with color C and thickness T
func (S *SDL) Line(x1, y1, x2, y2 int, C color.RGBA, T int) {
	var x, y float64
	x = float64(x2 - x1)
	y = float64(y2 - y1)
	len := math.Sqrt(float64(x*x + y*y))
	addx := x / len
	addy := y / len
	x = float64(x1)
	y = float64(y1)
	for i := 0; i < int(len); i++ {
		if int32(x) < S.w && int32(y) < S.h {
			S.Set(int(x), int(y), C, T)
		}
		x += addx
		y += addy
	}
}

func circleInit(R, x0, y0 float64) (float64, float64, float64, float64, float64) {
	err := -R
	x := R - 0.5
	y := 0.5
	cx := x0 - 0.5
	cy := y0 - 0.5
	return err, x, y, cx, cy
}

func correction(err, x, y float64) (float64, float64, float64) {
	err += y
	y++
	err += y
	if err >= 0 {
		x--
		err -= x
		err -= x
	}
	return err, x, y
}

// Circle Draw a circle of center (x0, y0) and radius R
// with color C and thickness T
func (S *SDL) Circle(x0, y0, R int, C color.RGBA, T int) {
	err, x, y, cx, cy := circleInit(float64(R), float64(x0), float64(y0))
	for {
		S.Set(int(cx+x), int(cy+y), C, T)
		S.Set(int(cx+y), int(cy+x), C, T)
		if x != 0 {
			S.Set(int(cx-x), int(cy+y), C, T)
			S.Set(int(cx+y), int(cy-x), C, T)
		}
		if y != 0 {
			S.Set(int(cx+x), int(cy-y), C, T)
			S.Set(int(cx-y), int(cy+x), C, T)
		}
		if x != 0 && y != 0 {
			S.Set(int(cx-x), int(cy-y), C, T)
			S.Set(int(cx-y), int(cy-x), C, T)
		}
		err, x, y = correction(err, x, y)
		if x < y {
			break
		}
	}
}

// Sector Draw a circle sector of center (x0, y0) and radius R
// with color C and thickness T
func (S *SDL) Sector(x0, y0, R int, C color.RGBA, T int) {
	err, x, y, cx, cy := circleInit(float64(R), float64(x0), float64(y0))
	for {
		if x != 0 {
			if cx-x > float64(x0) {
				S.Set(int(cx-x), int(cy+y), C, T)
			}
			S.Set(int(cx+y), int(cy-x), C, T)
		}
		if y != 0 {
			S.Set(int(cx+x), int(cy-y), C, T)
			if cx-y > float64(x0) {
				S.Set(int(cx-y), int(cy+x), C, T)
			}
		}
		err, x, y = correction(err, x, y)
		if x < y {
			break
		}
	}
}

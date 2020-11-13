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

// Refresh update surface
func (S *SDL) Refresh() {
	S.win.UpdateSurface()
}

type neighbor struct {
	x, y int
}

// SetBigPoint Set a big point at coordinates
func (S *SDL) SetBigPoint(x, y int, C color.RGBA) {
	var neighbors = []neighbor{
		{0, 1}, {1, 1}, {1, 0}, {1, -1},
		{0, -1}, {-1, -1}, {-1, 0}, {-1, 1},
	}
	S.sur.Set(x, y, C)
	for _, n := range neighbors {
		S.sur.Set(x+n.x, y+n.y, C)
	}
}

// Line Draw a line from (x1, y1) to (x2, y2)
func (S *SDL) Line(x1, y1, x2, y2 int, C color.RGBA) {
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
			S.sur.Set(int(x), int(y), C)
		}
		x += addx
		y += addy
	}
}

// Circle Draw a circle of center (x0, y0) and radius R
func (S *SDL) Circle(x0, y0, R int, C color.RGBA) {
	err := float64(-R)
	x := float64(R) - 0.5
	y := 0.5
	cx := float64(x0) - 0.5
	cy := float64(y0) - 0.5

	for {
		S.sur.Set(int(cx+x), int(cy+y), C)
		S.sur.Set(int(cx+y), int(cy+x), C)
		if x != 0 {
			S.sur.Set(int(cx-x), int(cy+y), C)
			S.sur.Set(int(cx+y), int(cy-x), C)
		}
		if y != 0 {
			S.sur.Set(int(cx+x), int(cy-y), C)
			S.sur.Set(int(cx-y), int(cy+x), C)
		}
		if x != 0 && y != 0 {
			S.sur.Set(int(cx-x), int(cy-y), C)
			S.sur.Set(int(cx-y), int(cy-x), C)
		}
		err += y
		y++
		err += y
		if err >= 0 {
			x--
			err -= x
			err -= x
		}
		if x < y {
			break
		}
	}
}

// Sector Draw a circle sector of center (x0, y0) and radius R
func (S *SDL) Sector(x0, y0, R int, C color.RGBA) {
	err := float64(-R)
	x := float64(R) - 0.5
	y := 0.5
	cx := float64(x0) - 0.5
	cy := float64(y0) - 0.5

	for {
		if x != 0 {
			if cx-x > float64(x0) {
				S.sur.Set(int(cx-x), int(cy+y), C)
			}
			S.sur.Set(int(cx+y), int(cy-x), C)
		}
		if y != 0 {
			S.sur.Set(int(cx+x), int(cy-y), C)
			if cx-y > float64(x0) {
				S.sur.Set(int(cx-y), int(cy+x), C)
			}
		}
		err += y
		y++
		err += y
		if err >= 0 {
			x--
			err -= x
			err -= x
		}
		if x < y {
			break
		}
	}
}

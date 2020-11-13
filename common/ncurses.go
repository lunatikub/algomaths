package common

import (
	"log"

	gc "github.com/gbin/goncurses"
)

// NCInit Initialize a ncurses window
func NCInit() *gc.Window {
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

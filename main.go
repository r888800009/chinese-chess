package main

import "github.com/veandco/go-sdl2/sdl"

var surface *sdl.Surface
var window *sdl.Window
const fps = 60

func display() {
	surface.FillRect(nil, 0)
	window.UpdateSurface()
}

func main() {
    var err error
	println("Start window")
	window, err = sdl.CreateWindow("", sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

    surface, err = window.GetSurface()
	if err != nil {
		panic(err)
	}

	running := true
	for running {
		display()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
	}
}

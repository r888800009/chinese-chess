package main

import "github.com/veandco/go-sdl2/sdl"

var surface *sdl.Surface
var window *sdl.Window
var render *sdl.Renderer
const fps = 60
const size = 60
const winW, winH = 800, 800

func paintBoard() {
    render.SetDrawColor(255, 255, 0, sdl.ALPHA_OPAQUE)
    var centerX, centerY = winW / 2, winH / 2

    // vertical
    for x := -4; x <= 4; x++ {
        render.DrawLine(int32(centerX + x * size),
                        int32(centerY + -4 * size),
                        int32(centerX + x * size),
                        int32(centerY + 0 * size))

        render.DrawLine(int32(centerX + x * size),
                        int32(centerY + 5 * size),
                        int32(centerX + x * size),
                        int32(centerY + 1 * size))
    }

    render.DrawLine(int32(centerX + -4 * size),
                    int32(centerY + -4 * size),
                    int32(centerX + -4 * size),
                    int32(centerY + 5 * size))

    render.DrawLine(int32(centerX + 4 * size),
                    int32(centerY + -4 * size),
                    int32(centerX + 4 * size),
                    int32(centerY + 5 * size))


    // horizonatl
    for y := -4; y <= 5; y++ {
        render.DrawLine(int32(centerX + -4 * size),
                        int32(centerY + y * size),
                        int32(centerX + 4 * size),
                        int32(centerY + y * size))
     }

     // X
     render.DrawLine(int32(centerX + -1 * size),
                     int32(centerY + 3 * size),
                     int32(centerX + 1 * size),
                     int32(centerY + 5 * size))

     render.DrawLine(int32(centerX + 1 * size),
                     int32(centerY + 3 * size),
                     int32(centerX + -1 * size),
                     int32(centerY + 5 * size))

     render.DrawLine(int32(centerX + -1 * size),
                     int32(centerY + -4 * size),
                     int32(centerX + 1 * size),
                     int32(centerY + -2 * size))

     render.DrawLine(int32(centerX + 1 * size),
                     int32(centerY + -4 * size),
                     int32(centerX + -1 * size),
                     int32(centerY + -2 * size))
}

func display() {

    render.SetDrawColor(0x00, 0x00, 0x00, sdl.ALPHA_OPAQUE)
    render.Clear()

    paintBoard()

    render.Present()
}

func main() {
    var err error
	println("Start window")


	window, render, err = sdl.CreateWindowAndRenderer(winW, winH, sdl.WINDOW_SHOWN)
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

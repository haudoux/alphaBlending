package picture

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

//Pixels each pixel of the screen
type Pixels struct {
	screen []byte
}

var winWidth, winHeigth int

func initScreen() *Pixels {
	pixels := Pixels{}
	pixels.screen = make([]byte, winWidth*winHeigth*4)
	pixels.resetScreen()
	return &pixels
}

func (pixels *Pixels) resetScreen() {
	for i := range pixels.screen {
		pixels.screen[i] = 0
	}
}

func startSDL2(name string) (*sdl.Window, *sdl.Renderer, *sdl.Texture) {

	window, err := sdl.CreateWindow(name, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(winWidth), int32(winHeigth), sdl.WINDOW_SHOWN)

	if err != nil {
		fmt.Println(err)
		return window, nil, nil
	}
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
		return window, renderer, nil
	}
	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(winWidth), int32(winHeigth))
	if err != nil {
		fmt.Println(err)
	}
	return window, renderer, texture
}

//Run Start the sdl
func Run(name string, width, height int) {
	winHeigth = height
	winWidth = width
	window, renderer, texture := startSDL2(name)
	defer sdl.Quit()
	defer window.Destroy()
	if renderer != nil {
		defer renderer.Destroy()
	}
	if texture != nil {
		defer texture.Destroy()
	}
	mainLoop(renderer, texture)

}

func mainLoop(renderer *sdl.Renderer, texture *sdl.Texture) {
	var frameStart time.Time
	var elapsedTime float32

	pixels := initScreen()
	balloonTextures := LoadPicture()
	dir := 1
	for {
		frameStart = time.Now()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		pixels.resetScreen()

		for _, txt := range balloonTextures {
			txt.DrawAlpha(*pixels)
		}

		balloonTextures[1].X += float32(1 * dir)
		if balloonTextures[1].X > 340 || balloonTextures[1].X < 0 {
			dir = dir * -1
		}
		balloonTextures[2].Y += float32(1 * dir)
		if balloonTextures[2].Y > 340 || balloonTextures[2].Y < 0 {
			dir = dir * -1
		}

		err := texture.Update(nil, pixels.screen, winWidth*4)
		if err != nil {
			fmt.Println(sdl.GetError())
		}
		err = renderer.Copy(texture, nil, nil)
		if err != nil {
			fmt.Println(sdl.GetError())
		}

		renderer.Present()

		elapsedTime = float32(time.Since(frameStart).Seconds())

		//144 FPS
		if elapsedTime < 0.0069 {
			sdl.Delay(5 - uint32(elapsedTime)*1000.0)
			elapsedTime = float32(time.Since(frameStart).Seconds())
		}
	}

}

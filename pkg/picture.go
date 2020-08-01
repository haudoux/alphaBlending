package picture

import (
	"image/png"
	"os"
)

//Texture texture type
type Texture struct {
	TexturePixels []byte
	W,H, Pitch int
	Pos
}

//Draw the ballon in the pixels
func (txt *Texture) Draw(pixels Pixels) Pixels{
	for y := 0; y < txt.H; y++ {
		for x := 0; x < txt.W; x++{
			screenY := y + int(txt.Y)
			screenX := x + int(txt.X)
			if(screenX >= 0 && screenX < winWidth && screenY >= 0 && screenY < winHeigth) {
				txtIndex := y * txt.Pitch + x * 4
				screenIndex := screenY * winWidth * 4 + screenX * 4 
				pixels.screen[screenIndex] = txt.TexturePixels[txtIndex]
				pixels.screen[screenIndex+1] = txt.TexturePixels[txtIndex+1]
				pixels.screen[screenIndex+2] = txt.TexturePixels[txtIndex+2]
				pixels.screen[screenIndex+3] = txt.TexturePixels[txtIndex+3]
			}
		}
	}
	return pixels
}

//DrawAlpha draw the ballon with alpha blending in the pixels
func (txt *Texture) DrawAlpha(pixels Pixels) Pixels{
	for y := 0; y < txt.H; y++ {
		screenY := y + int(txt.Y)
		if screenY >= 0 && screenY < winHeigth {
			for x := 0; x < txt.W; x++{
				screenX := x + int(txt.X)
				if screenX >= 0 && screenX < winWidth {
					txtIndex := y * txt.Pitch + x * 4
					screenIndex := screenY * winWidth * 4 + screenX * 4 

					srcR := int(txt.TexturePixels[txtIndex])
					srcG := int(txt.TexturePixels[txtIndex + 1])
					srcB := int(txt.TexturePixels[txtIndex + 2])
					srcA := int(txt.TexturePixels[txtIndex + 3])

					newR := int(pixels.screen[screenIndex])
					newG := int(pixels.screen[screenIndex + 1])
					newB := int(pixels.screen[screenIndex + 2])

					rstR := (srcR*255 + newR*(255-srcA)) / 255
					rstG := (srcG*255 + newG*(255-srcA)) / 255
					rstB := (srcB*255 + newB*(255-srcA)) / 255

					pixels.screen[screenIndex] = byte(rstR)
					pixels.screen[screenIndex + 1] = byte(rstG)
					pixels.screen[screenIndex + 2] = byte(rstB)
				}
			}	
		}

	}
	return pixels
}
//LoadPicture load a picture and return a pixels byte of the picture
func LoadPicture() []Texture{

	balloonsPath := []string{"balloon_blue.png","balloon_red.png","balloon_green.png"}
	ballonTexture := make([]Texture, len(balloonsPath))
	for i, path := range balloonsPath {
		file, err := os.Open("balloons/"+path)
		if err != nil {
			panic(err)
		}
		defer file.Close()
	
		img, err := png.Decode(file)
		if err != nil {
			panic(err)
		}
	
		width := img.Bounds().Max.X
		heigth := img.Bounds().Max.Y
	
		pixels := make([]byte,width*heigth*4)
		pIndex := 0
		for y := 0; y < heigth; y++{
			for x := 0; x < width; x++ {
				r,g,b,a := img.At(x,y).RGBA()
				pixels[pIndex] = byte(r / 256)
				pIndex++
				pixels[pIndex] = byte(g / 256)
				pIndex++
				pixels[pIndex] = byte(b / 256)
				pIndex++
				pixels[pIndex] = byte(a / 256)
				pIndex++
			}
		}
		ballonTexture[i] = Texture{pixels,width,heigth, width*4, Pos{0,0}}
	}
	
	return ballonTexture
}
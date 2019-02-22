package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil" // required for isKeyJustPressed
 	"github.com/hajimehoshi/ebiten/ebitenutil" // required for debug text
)

type gameState int

const (
	titleScreen gameState = iota
	play
)

var (
	state  gameState 
	square *ebiten.Image
)

func update(screen *ebiten.Image) error {

	screen.Fill(color.NRGBA{0x00, 0x00, 0x00, 0xff})

	if state == titleScreen {
	
		ebitenutil.DebugPrint(screen, "Title screen")
		if square == nil { 
			square, _ = ebiten.NewImage(32, 32, ebiten.FilterNearest)
		}
		someColor := &color.NRGBA{0x00, 0x80, 0x80, 0xff}
		square.Fill(someColor)

		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(64.0, 64.0)
		screen.DrawImage(square, opts)

		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			state = play
			return nil
		}

	}

	if state == play {
		ebitenutil.DebugPrint(screen, "Play screen")

		if square == nil { 
			square, _ = ebiten.NewImage(32, 32, ebiten.FilterNearest)
		}
		someColor := &color.NRGBA{0xff, 0xa5, 0x00, 0xff}
		square.Fill(someColor)

		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(64.0, 64.0)
		screen.DrawImage(square, opts)

		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			state = titleScreen
			return nil
		}

	}

	return nil
}

func main() {

	state = titleScreen

	if err := ebiten.Run(update, 400, 300, 2, "State!"); err != nil {
		panic(err)
	}
}

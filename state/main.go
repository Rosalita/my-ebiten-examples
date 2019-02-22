package main

import (
	"log"
	"golang.org/x/image/font"
	"image/color"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"


	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil" // required for isKeyJustPressed
 	"github.com/hajimehoshi/ebiten/ebitenutil" // required for debug text
)

type gameState int

const (
	titleScreen gameState = iota
	options
	play
)

var (
	state  gameState
	playButton *ebiten.Image 
	optionsButton *ebiten.Image
	square *ebiten.Image
	mplusNormalFont font.Face
	mplusBigFont    font.Face
)

func init() {
	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	mplusBigFont = truetype.NewFace(tt, &truetype.Options{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

func update(screen *ebiten.Image) error {

	screen.Fill(color.NRGBA{0x00, 0x00, 0x00, 0xff})

	if state == titleScreen {
	
	
		ebitenutil.DebugPrint(screen, "Title screen")
		if playButton == nil { 
			playButton, _ = ebiten.NewImage(128, 32, ebiten.FilterNearest)
		}

		if optionsButton == nil { 
			optionsButton, _ = ebiten.NewImage(128, 32, ebiten.FilterNearest)
		}
		someColor := &color.NRGBA{0x00, 0x80, 0x80, 0xff}
		playButton.Fill(someColor)
		optionsButton.Fill(someColor)

		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(128.0, 128.0)

		text.Draw(playButton, "PLAY", mplusNormalFont, 36, 25, color.White)

		screen.DrawImage(playButton, opts)

		opts.GeoM.Translate(0, 36.0)

		text.Draw(optionsButton, "OPTIONS", mplusNormalFont, 12, 25, color.White)
		screen.DrawImage(optionsButton, opts)
	



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

		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
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

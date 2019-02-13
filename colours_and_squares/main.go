package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil" // This is required to draw debug texts.
)

var (
	square    *ebiten.Image
	square2   *ebiten.Image
	square3   *ebiten.Image
	square4   *ebiten.Image
	square5   *ebiten.Image
	someColor *color.NRGBA
	opts      *ebiten.DrawImageOptions
)

func update(screen *ebiten.Image) error {
	// NRGBA represents a non-alpha-premultiplied 32-bit color.
	screen.Fill(color.NRGBA{0xff, 0xcc, 0xf9, 0xff})
	ebitenutil.DebugPrint(screen, "Colours and Squares!")

	if square == nil { // if square isn't set to an image
		// set it to a new image
		square, _ = ebiten.NewImage(32, 32, ebiten.FilterNearest)
	}

	if square2 == nil { // if square isn't set to an image
		// set it to a new image
		square2, _ = ebiten.NewImage(64, 64, ebiten.FilterNearest)
	}

	if square3 == nil { // if square isn't set to an image
		// set it to a new image
		square3, _ = ebiten.NewImage(128, 128, ebiten.FilterNearest)
	}

	if square4 == nil { // if square isn't set to an image
		// set it to a new image
		square4, _ = ebiten.NewImage(256, 256, ebiten.FilterNearest)
	}

	if square5 == nil { // if square isn't set to an image
	// set it to a new image
	square5, _ = ebiten.NewImage(512, 512, ebiten.FilterNearest)
}
	// set the colour of the squares
	square.Fill(someColor)
	square2.Fill(someColor)
	square3.Fill(someColor)
	square4.Fill(someColor)
	square5.Fill(someColor)

	someColor.R++
	someColor.G--

	rand.Seed(time.Now().UnixNano())
	offset := rand.Intn(64) - 32

	opts.GeoM.Translate(float64(offset), float64(offset))

	// create render options that tell Ebiten how to draw image to screen
	// setting a Geometry matrix in the options allows shapes to be
	// translated, enlarged or rotated

	// translate(tx, ty, float64)
	// tx is the distance from the left, also called x offset
	// ty is the distance from the right, also called y offset

	screen.DrawImage(square, opts)
	screen.DrawImage(square2, opts)
	screen.DrawImage(square3, opts)
	screen.DrawImage(square4, opts)
	screen.DrawImage(square5, opts)

	return nil
}

func main() {
	someColor = &color.NRGBA{0xff, 0xaf, 0xed, 0x55}
	opts = &ebiten.DrawImageOptions{}

	if err := ebiten.Run(update, 320, 240, 2, "Colours and Squares!"); err != nil {
		panic(err)
	}
}

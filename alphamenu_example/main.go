package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil" // This is required to draw debug texts.
	"github.com/hajimehoshi/ebiten/inpututil"  // required for isKeyJustPressed

	am "github.com/Rosalita/my-ebiten/pkgs/alphamenu"

)

var (
	alphaMenu am.AlphaMenu
)

func update(screen *ebiten.Image) error {

	screen.Fill(color.NRGBA{0x00, 0x00, 0x00, 0xff})
	ebitenutil.DebugPrint(screen, "Alphabet menu")

	alphaMenu.Draw(screen)

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		alphaMenu.DecY()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		alphaMenu.IncY()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		alphaMenu.IncX()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		alphaMenu.DecX()
	}

	return nil
}

func main() {

	white := &color.NRGBA{0xff, 0xff, 0xff, 0xff}

	alphaMenuInput := am.Input{
		Tx: 50,
		Ty: 50,
		DefaultBgColour: white,
	}

	alphaMenu, _ = am.NewMenu(alphaMenuInput)


	if err := ebiten.Run(update, 320, 240, 2, "Alphabet menu"); err != nil {
		panic(err)
	}
}

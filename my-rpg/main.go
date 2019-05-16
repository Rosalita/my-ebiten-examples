package main

import (
	"bytes"
	"image"
	"image/color"
	"log"
	"os"

	im "github.com/Rosalita/ebiten-pkgs/imagemenu"
	lm "github.com/Rosalita/ebiten-pkgs/listmenu"
	"github.com/Rosalita/my-ebiten-examples/resources/my_img"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil" // required for debug text
	"github.com/hajimehoshi/ebiten/inpututil"  // required for isKeyJustPressed
)

type gameState int

const (
	titleScreen gameState = iota
	options
	charCreation
	quit
)

// define some kind of palette
var (
	white   = &color.NRGBA{0xff, 0xff, 0xff, 0xff}
	pink    = &color.NRGBA{0xff, 0x69, 0xb4, 0xff}
	orange1 = &color.NRGBA{0xfe, 0x7f, 0x2d, 0xff}
	blue1   = &color.NRGBA{0x6f, 0xe9, 0xee, 0xff}
	green1  = &color.NRGBA{0x00, 0x38, 0x40, 0xff}
	green2  = &color.NRGBA{0x00, 0x5a, 0x5b, 0xff}
	green3  = &color.NRGBA{0x00, 0x73, 0x69, 0xff}
	green4  = &color.NRGBA{0x00, 0x8c, 0x72, 0xff}
	green5  = &color.NRGBA{0x02, 0xa6, 0x76, 0xff}
	purple1 = &color.NRGBA{0x30, 0x28, 0x40, 0xff}
	purple2 = &color.NRGBA{0x47, 0x39, 0x5b, 0xff}
	purple3 = &color.NRGBA{0x5f, 0x49, 0x73, 0xff}
	purple4 = &color.NRGBA{0x7b, 0x58, 0x8c, 0xff}
	purple5 = &color.NRGBA{0x99, 0x69, 0xa6, 0xff}
)

var (
	state       gameState
	mainImage   *ebiten.Image
	charImage   *ebiten.Image
	rightArrow  *ebiten.Image
	leftArrow   *ebiten.Image
	mainMenu    lm.ListMenu
	optionsMenu lm.ListMenu
	charGroupMenu    im.ImageMenu
	humanMenu    im.ImageMenu
	creatureMenu im.ImageMenu
)

func init() {

	img, _, err := image.Decode(bytes.NewReader(my_img.BirdSkull))
	if err != nil {
		log.Fatal(err)
	}
	mainImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	rightArrow, _ = ebiten.NewImage(36, 36, ebiten.FilterDefault)
	leftArrow, _ = ebiten.NewImage(36, 36, ebiten.FilterDefault)

}

func update(screen *ebiten.Image) error {

	screen.Fill(color.NRGBA{0x00, 0x00, 0x00, 0xff})

	if state == titleScreen {

		ebitenutil.DebugPrint(screen, "Title screen")
		mainMenu.Draw(screen)

		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(200, 24)
		screen.DrawImage(mainImage, opts)

		if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
			mainMenu.DecrementSelected()
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
			mainMenu.IncrementSelected()
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			switch mainMenu.GetSelectedItem() {
			case "playButton":
				state = charCreation
			case "optionButton":
				state = options
			case "quitButton":
				os.Exit(0)
			}
			return nil
		}

	}

	if state == charCreation {
		ebitenutil.DebugPrint(screen, "Character Creation")

		charGroupMenu.Draw(screen)

		humanMenu.Draw(screen)

		creatureMenu.Draw(screen)

		if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
			charGroupMenu.IncrementSelected()
			humanMenu.IncrementSelected()
			creatureMenu.IncrementSelected()
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
			charGroupMenu.DecrementSelected()
			humanMenu.DecrementSelected()
			creatureMenu.DecrementSelected()
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			state = titleScreen
			return nil
		}

	}

	if state == options {
		ebitenutil.DebugPrint(screen, "Options screen")
		optionsMenu.Draw(screen)

		if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
			optionsMenu.DecrementSelected()
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
			optionsMenu.IncrementSelected()
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			state = titleScreen
			return nil
		}
	}

	return nil
}

func main() {

	initMenus()

	state = titleScreen

	if err := ebiten.Run(update, 400, 300, 2, "State!"); err != nil {
		panic(err)
	}
}

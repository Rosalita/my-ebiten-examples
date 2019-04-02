package main

import (
	"bytes"
	"image"
	"image/color"
	"log"
	"os"

	menu "github.com/Rosalita/my-ebiten/pkgs/listmenu"
	"github.com/Rosalita/my-ebiten/resources/avatars"
	"github.com/Rosalita/my-ebiten/resources/my_img"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil" // required for debug text
	"github.com/hajimehoshi/ebiten/inpututil"  // required for isKeyJustPressed
)

type gameState int

const (
	titleScreen gameState = iota
	options
	charSel
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
	mainMenu    menu.ListMenu
	optionsMenu menu.ListMenu
)

func init() {

	// use https://github.com/hajimehoshi/file2byteslice
	// to embed an image

	img, _, err := image.Decode(bytes.NewReader(my_img.BirdSkull))
	if err != nil {
		log.Fatal(err)
	}
	mainImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	img2, _, err := image.Decode(bytes.NewReader(avatars.F_02_s))
	if err != nil {
		log.Fatal(err)
	}

	charImage, _ = ebiten.NewImageFromImage(img2, ebiten.FilterDefault)

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
				state = charSel
			case "optionButton":
				state = options
			case "quitButton":
				os.Exit(0)
			}
			return nil
		}

	}

	if state == charSel {
		ebitenutil.DebugPrint(screen, "Character Select")

		picx := 100.0
		picy := 24.0

		rv1 := ebiten.Vertex{
			DstX:   0 + float32(picx*2),
			DstY:   0 + float32(picy*2),
			SrcX:   0,
			SrcY:   0,
			ColorR: 1,
			ColorG: 1,
			ColorB: 1,
			ColorA: 1,
		}

		rv2 := ebiten.Vertex{
			DstX:   10 + float32(picx*2),
			DstY:   10 + float32(picy*2),
			SrcX:   0,
			SrcY:   0,
			ColorR: 1,
			ColorG: 1,
			ColorB: 1,
			ColorA: 1,
		}

		rv3 := ebiten.Vertex{
			DstX:   0 + float32(picx*2),
			DstY:   20 + float32(picy*2),
			SrcX:   0,
			SrcY:   0,
			ColorR: 1,
			ColorG: 1,
			ColorB: 1,
			ColorA: 1,
		}

		lv1 := ebiten.Vertex{
			DstX:   10 + float32(picx),
			DstY:   0 + float32(picy*2),
			SrcX:   0,
			SrcY:   0,
			ColorR: 1,
			ColorG: 1,
			ColorB: 1,
			ColorA: 1,
		}

		lv2 := ebiten.Vertex{
			DstX:   10 + float32(picx),
			DstY:   20 + float32(picy*2),
			SrcX:   0,
			SrcY:   0,
			ColorR: 1,
			ColorG: 1,
			ColorB: 1,
			ColorA: 1,
		}

		lv3 := ebiten.Vertex{
			DstX:   0 + float32(picx),
			DstY:   10 + float32(picy*2),
			SrcX:   0,
			SrcY:   0,
			ColorR: 1,
			ColorG: 1,
			ColorB: 1,
			ColorA: 1,
		}

		rightvs := []ebiten.Vertex{rv1, rv2, rv3}
		leftvs := []ebiten.Vertex{lv1, lv2, lv3}

		indices := []uint16{0, 1, 2}

		op := &ebiten.DrawTrianglesOptions{}

		rightArrow.Fill(color.White)
		leftArrow.Fill(color.White)

		screen.DrawTriangles(rightvs, indices, rightArrow, op)
		screen.DrawTriangles(leftvs, indices, leftArrow, op)

		// if charImage == nil {
		// 	charImage, _ = ebiten.NewImage(100, 100, ebiten.FilterNearest)
		// }

		// charImage.Fill(color.White)

		// opts := &ebiten.DrawImageOptions{}
		// opts.GeoM.Translate(150, 24)

		// screen.DrawImage(charImage, opts)

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Translate(picx, picy)
		screen.DrawImage(charImage, opts)

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

func initMenus() {

	mainMenuItems := []menu.Item{
		{Name: "playButton",
			Text:     "PLAY",
			TxtX:     40,
			TxtY:     25,
			BgColour: white},
		{Name: "optionButton",
			Text:     "OPTIONS",
			TxtX:     16,
			TxtY:     25,
			BgColour: white},
		{Name: "quitButton",
			Text:     "QUIT",
			TxtX:     40,
			TxtY:     25,
			BgColour: white},
	}

	mainMenuInput := menu.Input{
		Width:              140,
		Height:             36,
		Tx:                 24,
		Ty:                 24,
		Offy:               40,
		DefaultSelBGColour: pink,
		Items:              mainMenuItems,
	}

	mainMenu, _ = menu.NewMenu(mainMenuInput)

	optionsMenuItems := []menu.Item{
		{Name: "screen",
			Text:     "SCREEN",
			TxtX:     28,
			TxtY:     25,
			BgColour: white},
		{Name: "sound",
			Text:     "SOUND",
			TxtX:     32,
			TxtY:     25,
			BgColour: white},
		{Name: "language",
			Text:     "LANGUAGE",
			TxtX:     4,
			TxtY:     25,
			BgColour: white},
	}

	optionsMenuInput := menu.Input{
		Width:              140,
		Height:             36,
		Tx:                 24,
		Ty:                 24,
		Offy:               40,
		DefaultSelBGColour: pink,
		Items:              optionsMenuItems,
	}

	optionsMenu, _ = menu.NewMenu(optionsMenuInput)

}

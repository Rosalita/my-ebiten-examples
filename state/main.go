package main

import (
	"image/color"
	"log"
	"os"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil" // required for debug text
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/inpututil" // required for isKeyJustPressed
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

type gameState int

const (
	titleScreen gameState = iota
	options
	play
	quit
)

type menuItem struct {
	name  string
	image *ebiten.Image
	text  string
}

// MenuList is a menu
type MenuList struct {
	Tx             float64      // x translation of the menu
	Ty             float64      // y translation of the menu
	Offx           float64      // x offset of subsequent menu items
	Offy           float64      // y offset of subsequent menu items
	BaseColour     *color.NRGBA // default unselected colour
	SelectedColour *color.NRGBA // colour used when button is selected
	SelectedIndex  *int         // index of the item in list which is selected
	MenuItems      []menuItem   // menu items
}

//NewMenu creates a new menu
func NewMenu(tx float64, ty float64, offx float64, offy float64, basecolour *color.NRGBA, selectedColour *color.NRGBA, MenuItems []menuItem) MenuList {
	defaultSelectedIndex := 0

	for i := range MenuItems {
		newImage, _ := ebiten.NewImage(128, 32, ebiten.FilterNearest)
		MenuItems[i].image = newImage
	}

	ml := MenuList{
		Tx:             tx,
		Ty:             ty,
		Offx:           offx,
		Offy:           offy,
		BaseColour:     basecolour,
		SelectedColour: selectedColour,
		SelectedIndex:  &defaultSelectedIndex,
		MenuItems:      MenuItems,
	}

	return ml
}

//GetBaseColour returns the menu base colour
func (m *MenuList) GetBaseColour() *color.NRGBA {
	return m.BaseColour
}

//GetSelectedColour returns the menu selected colour
func (m *MenuList) GetSelectedColour() *color.NRGBA {
	return m.SelectedColour
}

//GetSelectedItem returns then name of the selected item
func (m *MenuList) GetSelectedItem() string {
	return m.MenuItems[*m.SelectedIndex].name
}

//IncrementSelected increments the selected index provided it is not already at maximum
func (m *MenuList) IncrementSelected() {
	maxIndex := len(m.MenuItems) - 1
	if *m.SelectedIndex < maxIndex {
		*m.SelectedIndex++
	}
}

//DecrementSelected decrements the selected index provided it is not already at minimum
func (m *MenuList) DecrementSelected() {
	minIndex := 0
	if *m.SelectedIndex > minIndex {
		*m.SelectedIndex--
	}
}

//Draw draws the menu to the screen
func (m *MenuList) Draw(screen *ebiten.Image) {

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(m.Tx, m.Ty)

	for index, item := range m.MenuItems {

		if index == *m.SelectedIndex {
			item.image.Fill(m.GetSelectedColour())
		} else {
			item.image.Fill(m.GetBaseColour())
		}

		textX := 0
		if len(item.text) == 4 {
			textX = 36
		}
		if len(item.text) == 7 {
			textX = 12
		}

		text.Draw(item.image, item.text, mplusNormalFont, textX, 25, color.White)
		screen.DrawImage(item.image, opts)
		opts.GeoM.Translate(m.Offx, m.Offy)
	}
}

var (
	state           gameState
	playImage       *ebiten.Image
	optionsImage    *ebiten.Image
	quitImage       *ebiten.Image
	square          *ebiten.Image
	mplusNormalFont font.Face
	mplusBigFont    font.Face
	mainMenu        MenuList
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

}

func update(screen *ebiten.Image) error {

	screen.Fill(color.NRGBA{0x00, 0x00, 0x00, 0xff})

	if state == titleScreen {

		ebitenutil.DebugPrint(screen, "Title screen")
		mainMenu.Draw(screen)

		if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
			mainMenu.DecrementSelected()
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
			mainMenu.IncrementSelected()
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			switch mainMenu.GetSelectedItem() {
			case "playButton":
				state = play
			case "optionButton":
				state = options
			case "quitButton":
				os.Exit(0)
			}
			return nil
		}

	}

	if state == play {
		ebitenutil.DebugPrint(screen, "Play screen")

		if square == nil {
			square, _ = ebiten.NewImage(32, 32, ebiten.FilterNearest)
		}
		someColor := &color.NRGBA{0x7f, 0xff, 0x00, 0xff}
		square.Fill(someColor)

		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(64.0, 64.0)
		screen.DrawImage(square, opts)

		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			state = titleScreen
			return nil
		}

	}

	if state == options {
		ebitenutil.DebugPrint(screen, "Options screen")

		if square == nil {
			square, _ = ebiten.NewImage(32, 32, ebiten.FilterNearest)
		}
		someColor := &color.NRGBA{0x8a, 0x2b, 0xe2, 0xff}
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

	menuItems := []menuItem{
		{name: "playButton", text: "PLAY"},
		{name: "optionButton", text: "OPTIONS"},
		{name: "quitButton", text: "QUIT"},
	}

	mainMenu = NewMenu(128, 128, 0, 30, &color.NRGBA{0x00, 0x80, 0x80, 0xff}, &color.NRGBA{0xff, 0xa5, 0x00, 0xff}, menuItems)

	state = titleScreen

	if err := ebiten.Run(update, 400, 300, 2, "State!"); err != nil {
		panic(err)
	}
}

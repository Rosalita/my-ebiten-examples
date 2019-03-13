package main

import (
	"errors"
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

// MenuItem represents an item in a menu list
type MenuItem struct {
	Name  string
	image *ebiten.Image
	Text  string
}

// MenuList is a navigatable, selectable menu
type MenuList struct {
	Tx             float64      // x translation of the menu
	Ty             float64      // y translation of the menu
	Width          int          // width of all menu items
	Height         int          // height of all menu items
	Offx           float64      // x offset of subsequent menu items
	Offy           float64      // y offset of subsequent menu items
	BaseColour     *color.NRGBA // default unselected colour
	SelectedColour *color.NRGBA // colour used when button is selected
	SelectedIndex  *int         // index of the item in list which is selected
	MenuItems      []MenuItem   // menu items
}

// MenuListInput is an object used to create a menu list
type MenuListInput struct {
	Tx             float64      // optional, x translation of the menu, if not provided will be 0
	Ty             float64      // optional, y translation of the menu, if not provided will be 0
	Width          int          // mandatory, width of all menu items
	Height         int          // mandatory, height of all menu items
	Offx           float64      // optional, offset of subsequent menu items, if not provided will 0
	Offy           float64      // optional, offset of subsequent menu items, if not provided will be menu item height
	BaseColour     *color.NRGBA // optional, base colour of menu, if not provided will be cyan
	SelectedColour *color.NRGBA // optional, selected colour of menu, if not provided will be magenta
	MenuItems      []MenuItem   // mandtory, list of menu items
}

//NewMenu constructs a new menu from a MenuListInput
func NewMenu(input MenuListInput) (MenuList, error) {

	if input.Width == 0 {
		return MenuList{}, errors.New("Mandatory input field width is missing")
	}
	if input.Height == 0 {
		return MenuList{}, errors.New("Mandatory input field height is missing")
	}
	if len(input.MenuItems) < 1 {
		return MenuList{}, errors.New("Mandatory input field MenuItems is missing")
	}

	if input.Offy == 0 {
		input.Offy = float64(input.Height)
	}

	if input.BaseColour == nil {
		input.BaseColour = &color.NRGBA{0x00, 0xff, 0xff, 0xff}
	}

	if input.SelectedColour == nil {
		input.SelectedColour = &color.NRGBA{0xff, 0x00, 0xff, 0xff}
	}

	defaultSelectedIndex := 0

	ml := MenuList{
		Tx:             input.Tx,
		Ty:             input.Ty,
		Width:          input.Width,
		Height:         input.Height,
		Offx:           input.Offx,
		Offy:           input.Offy,
		BaseColour:     input.BaseColour,
		SelectedColour: input.SelectedColour,
		SelectedIndex:  &defaultSelectedIndex,
		MenuItems:      input.MenuItems,
	}

	// initialise images for each menu item
	for i := range ml.MenuItems {
		newImage, _ := ebiten.NewImage(ml.Width, ml.Height, ebiten.FilterNearest)
		ml.MenuItems[i].image = newImage
	}

	return ml, nil
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
	return m.MenuItems[*m.SelectedIndex].Name
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
		if len(item.Text) == 4 {
			textX = 36
		}
		if len(item.Text) == 7 {
			textX = 12
		}

		text.Draw(item.image, item.Text, mplusNormalFont, textX, 25, color.White)
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

	newMenuItems := []MenuItem{
		{Name: "playButton", Text: "PLAY"},
		{Name: "optionButton", Text: "OPTIONS"},
		{Name: "quitButton", Text: "QUIT"},
	}

	newMenuInput := MenuListInput{
		Width:          128,
		Height:         36,
		Tx:             128,
		Ty:             128,
		BaseColour:     &color.NRGBA{0x00, 0x80, 0x80, 0xff},
		SelectedColour: &color.NRGBA{0xff, 0xa5, 0x00, 0xff},
		MenuItems:      newMenuItems,
	}

	newMenu, err := NewMenu(newMenuInput)

	if err != nil {
		log.Printf("unable to create menu: %+v\n", err)
	}

	mainMenu = newMenu

	state = titleScreen

	if err := ebiten.Run(update, 400, 300, 2, "State!"); err != nil {
		panic(err)
	}
}

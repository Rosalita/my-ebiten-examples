package alphamenu

import (
	"fmt"
	"image/color"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

var mplusNormalFont font.Face

func init() {
	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    16,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

// CharBox is a selectable box that holds a single character
type CharBox struct {
	Name   string        // a name to describe each menu item
	Char   string        // The character displayed on the box
	LocX   float64       // X location of the charbox
	LocY   float64       // Y location of the charbox
	TxtX   int           // X location to draw text
	TxtY   int           // Y location to draw text
	Xindex int           // the X index of the charbox
	Yindex int           // the Y index of the charbox
	image  *ebiten.Image // used to store the image for the char box
}

// AlphaMenu is a navigatable, selectable menu
type AlphaMenu struct {
	Tx                  float64      // x translation of the menu
	Ty                  float64      // y translation of the menu
	Offx                float64      // X offset for each character box
	Offy                float64      // Y offset for each character box
	DefaultBgColour     *color.NRGBA // default background colour
	DefaultTxtColour    *color.NRGBA // default text colour
	DefaultSelBgColour  *color.NRGBA // default selected background colour
	DefaultSelTxtColour *color.NRGBA // default selected text colour
	SelectedX           *int         // x index of the selected item
	SelectedY           *int         // y index of the selected item
	CharBoxes           []CharBox
}

// Input is an object used to create an alpha list
type Input struct {
	Tx                  float64      // optional, x translation of the menu, if not provided will be 0
	Ty                  float64      // optional, y translation of the menu, if not provided will be 0
	DefaultBgColour     *color.NRGBA // optional, default background colour of menu, if not provided will be cyan
	DefaultTxtColour    *color.NRGBA // optional, default text colour, if not provided will be black
	DefaultSelBGColour  *color.NRGBA // optional, default selected background colour of menu, if not provided will be magenta
	DefaultSelTxtColour *color.NRGBA //optional, default selected text colour of menu, if not provided it will be white
}

//NewMenu constructs a new alpha menu from a Input
func NewMenu(input Input) (AlphaMenu, error) {

	if input.DefaultBgColour == nil {
		input.DefaultBgColour = &color.NRGBA{0x00, 0xff, 0xff, 0xff}
	}

	if input.DefaultTxtColour == nil {
		input.DefaultTxtColour = &color.NRGBA{0x00, 0x00, 0x00, 0xff}
	}

	if input.DefaultSelBGColour == nil {
		input.DefaultSelBGColour = &color.NRGBA{0xff, 0x00, 0xff, 0xff}
	}

	if input.DefaultSelTxtColour == nil {
		input.DefaultSelTxtColour = &color.NRGBA{0xff, 0xff, 0xff, 0xff}
	}

	defaultSelectedX := 0
	defaultSelectedY := 0
	defaultOffx := 24.0
	defaultOffy := 24.0


	// draw 13 x 2 boxes for each lower case letter



	validChars := "abcdef"

	allBoxes := []CharBox{}


	for i, char := range validChars {

		if i == 0 {
			img, _ := ebiten.NewImage(20, 20, ebiten.FilterNearest)

			charBox := CharBox{
				Name:   string(char),
				Char:   string(char),
				LocX:   input.Tx,
				LocY:   input.Ty,
				TxtX:   6,
				TxtY:   14,
				Xindex: i,
				Yindex: 0,
				image:  img,
			}

			allBoxes = append(allBoxes, charBox)
			continue
		}

		img, _ := ebiten.NewImage(20, 20, ebiten.FilterNearest)

		charBox := CharBox{
			Name:   string(char),
			Char:   string(char),
			LocX:   allBoxes[i-1].LocX + float64(defaultOffx),
			LocY:   allBoxes[i-1].LocY,
			TxtX:   6,
			TxtY:   14,
			Xindex: i,
			Yindex: 0,
			image:  img,
		}

		allBoxes = append(allBoxes, charBox)

	}

	m := AlphaMenu{
		Tx:                  input.Tx,
		Ty:                  input.Ty,
		Offx:                defaultOffx,
		Offy:                defaultOffy,
		DefaultBgColour:     input.DefaultBgColour,
		DefaultTxtColour:    input.DefaultTxtColour,
		DefaultSelBgColour:  input.DefaultSelBGColour,
		DefaultSelTxtColour: input.DefaultSelTxtColour,
		SelectedX:           &defaultSelectedX,
		SelectedY:           &defaultSelectedY,
		CharBoxes:           allBoxes,
	}

	fmt.Println(m)

	return m, nil
}

//Draw draws the list menu to the screen
func (m *AlphaMenu) Draw(screen *ebiten.Image) {

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(m.Tx, m.Ty)

	for _, cb := range m.CharBoxes {

		cb.image.Fill(m.DefaultBgColour)

		text.Draw(cb.image, cb.Char, mplusNormalFont, int(cb.TxtX), int(cb.TxtY), m.DefaultTxtColour)

		screen.DrawImage(cb.image, opts)
		opts.GeoM.Translate(m.Offx, 0.0)
	}
}

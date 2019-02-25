package main

import (
	"image/color"
	"log"
	"fmt"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil" // required for debug text
	"github.com/hajimehoshi/ebiten/inpututil"  // required for isKeyJustPressed
)

type gameState int

const (
	titleScreen gameState = iota
	options
	play
)

type button struct {
	name        string
	buttonImage *ebiten.Image
	buttonText  string
}

type ButtonList struct {
	BaseColour     *color.NRGBA // default unselected colour
	SelectedColour *color.NRGBA // colour used when button is selected
	SelectedIndex *int // item in list which is selected
	Buttons        []button
}

func (b *ButtonList) GetBaseColour()*color.NRGBA{
	return b.BaseColour
}

func (b *ButtonList) GetSelectedColour()*color.NRGBA{
	return b.SelectedColour
}

func(b *ButtonList) IncrementSelected(){

	fmt.Println("incremented")
	maxIndex := len(b.Buttons) - 1
	fmt.Println(maxIndex)
	fmt.Println(*b.SelectedIndex)
	if *b.SelectedIndex < maxIndex{
		fmt.Println("++")
		*b.SelectedIndex ++
		fmt.Println(*b.SelectedIndex)
	}
}

func(b *ButtonList) DecrementSelected(){

	fmt.Println("decremented")

	minIndex := 0
	if *b.SelectedIndex > minIndex{
		*b.SelectedIndex -- 
	}
}

func (b *ButtonList) Draw(screen *ebiten.Image) {

	opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(128.0, 128.0)

	for index, button := range b.Buttons {

		if index == *b.SelectedIndex{
			button.buttonImage.Fill(b.GetSelectedColour())
		} else{
			button.buttonImage.Fill(b.GetBaseColour())
		}
		
		textX := 0
		if len(button.buttonText) == 4{
			textX = 36
		}
		if len(button.buttonText) == 7 {
			textX = 12
		}

		text.Draw(button.buttonImage, button.buttonText, mplusNormalFont, textX, 25, color.White)
	    screen.DrawImage(button.buttonImage, opts)
		opts.GeoM.Translate(0, 36.0)

	}
}

var (
	state              gameState
	playButtonImage    *ebiten.Image
	optionsButtonImage *ebiten.Image
	square             *ebiten.Image
	mplusNormalFont    font.Face
	mplusBigFont       font.Face
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
		if playButtonImage == nil {
			playButtonImage, _ = ebiten.NewImage(128, 32, ebiten.FilterNearest)
		}

		if optionsButtonImage == nil {
			optionsButtonImage, _ = ebiten.NewImage(128, 32, ebiten.FilterNearest)
		}

		selectedIndex := 0

		buttonList := ButtonList{
			BaseColour: &color.NRGBA{0x00, 0x80, 0x80, 0xff},
			SelectedColour: &color.NRGBA{0xff, 0xa5, 0x00, 0xff},
			SelectedIndex: &selectedIndex,
			Buttons: []button{
				button{
					name:        "playButton",
					buttonImage: playButtonImage,
					buttonText:  "PLAY",
				},
				button{
					name:        "optionButton",
					buttonImage: optionsButtonImage,
					buttonText:  "OPTIONS",
				},
			},
		}

		buttonList.Draw(screen)

		if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
			buttonList.DecrementSelected()
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyDown){
			buttonList.IncrementSelected()
		}

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

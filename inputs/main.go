package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func update(screen *ebiten.Image) error {
	// Get the x, y position of the cursor from the CursorPosition() function
	x, y := ebiten.CursorPosition()

	// Display the information with "X: xx, Y: xx" format
	ebitenutil.DebugPrint(screen, fmt.Sprintf("X: %d, Y: %d", x, y))

	// When the "left mouse button" is pressed...
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		ebitenutil.DebugPrint(screen, "\nYou're pressing the 'LEFT' mouse button.")
	}
	// When the "right mouse button" is pressed...
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		ebitenutil.DebugPrint(screen, "\n\nYou're pressing the 'RIGHT' mouse button.")
	}
	// When the "middle mouse button" is pressed...
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle) {
		ebitenutil.DebugPrint(screen, "\n\nYou're pressing the 'MIDDLE' mouse button.")
	}
	// When the "up arrow key" is pressed..
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		ebitenutil.DebugPrint(screen, "\nYou're pressing the 'UP' button.")
	}
	// When the "down arrow key" is pressed..
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		ebitenutil.DebugPrint(screen, "\n\nYou're pressing the 'DOWN' button.")
	}
	// When the "left arrow key" is pressed..
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		ebitenutil.DebugPrint(screen, "\n\n\nYou're pressing the 'LEFT' button.")
	}
	// When the "right arrow key" is pressed..
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		ebitenutil.DebugPrint(screen, "\n\n\n\nYou're pressing the 'RIGHT' button.")
	}
	return nil
}

func main() {
	if err := ebiten.Run(update, 320, 240, 2, "Hello world!"); err != nil {
		panic(err)
	}
}

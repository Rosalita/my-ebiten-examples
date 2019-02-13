package main

import (
    "github.com/hajimehoshi/ebiten"
    "github.com/hajimehoshi/ebiten/ebitenutil" // This is required to draw debug texts.
)

func update(screen *ebiten.Image) error {
    ebitenutil.DebugPrint(screen, "Hello world, this is Ebiten!")
    return nil
}

func main() {
    if err := ebiten.Run(update, 320, 240, 2, "Hello world!"); err != nil {
        panic(err)
    }
}
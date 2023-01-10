package main

import (
	"bombom/internal/desktop"

	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(desktop.NewGame()); err != nil {
		log.Fatal(err)
	}
}

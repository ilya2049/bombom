package main

import (
	"bombom/internal/desktop"
	"fmt"

	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	theGame, err := desktop.NewGame()
	if err != nil {
		fmt.Println(fmt.Errorf("failed to initialize the game: %w", err))

		return
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(theGame); err != nil {
		log.Fatal(err)
	}
}

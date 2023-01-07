package main

import (
	"github.com/gonutz/prototype/draw"

	"bombom/internal/desktop"
)

func main() {
	if err := draw.RunWindow(
		"Bom-bom!",
		640,
		640,
		desktop.DrawGameIteration,
	); err != nil {
		panic(err)
	}
}

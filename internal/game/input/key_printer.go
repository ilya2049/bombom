package input

import (
	"context"
	"fmt"
)

type keyPrinter struct {
}

func (p *keyPrinter) handleKeyPressed(_ context.Context, e KeyPressed) error {
	fmt.Println("key pressed: ", e.Key)

	return nil
}

func (p *keyPrinter) handleKeyReleased(_ context.Context, e KeyReleased) error {
	fmt.Println("key released: ", e.Key)

	return nil
}

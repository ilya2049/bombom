package input

import (
	"context"
)

type crossKeyFilter struct {
	keyHandler

	pressedKey Key

	crossKeys map[Key]bool
}

func newCrossKeyFilter(nextHandler keyHandler) *crossKeyFilter {
	return &crossKeyFilter{
		keyHandler: nextHandler,
		crossKeys: map[Key]bool{
			KeyUp:    true,
			KeyDown:  true,
			KeyLeft:  true,
			KeyRight: true,
		},
	}
}

func (f *crossKeyFilter) handleKeyPressed(ctx context.Context, e KeyPressed) error {
	if ok := f.crossKeys[e.Key]; !ok {
		return f.keyHandler.handleKeyPressed(ctx, e)
	}

	if f.pressedKey == KeyUnknown {
		f.pressedKey = e.Key

		return f.keyHandler.handleKeyPressed(ctx, e)
	}

	return nil
}

func (f *crossKeyFilter) handleKeyReleased(ctx context.Context, e KeyReleased) error {
	if ok := f.crossKeys[e.Key]; !ok {
		return f.keyHandler.handleKeyReleased(ctx, e)
	}

	if f.pressedKey == e.Key {
		f.pressedKey = KeyUnknown

		return f.keyHandler.handleKeyReleased(ctx, e)
	}

	return nil
}

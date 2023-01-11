package input

import (
	"context"
)

type crossKeyDriver struct {
	keyHandler

	pressedKey     Key
	nextPressedKey Key

	crossKeys map[Key]bool
}

func newCrossKeyDriver(nextHandler keyHandler) *crossKeyDriver {
	return &crossKeyDriver{
		keyHandler: nextHandler,
		crossKeys: map[Key]bool{
			KeyUp:    true,
			KeyDown:  true,
			KeyLeft:  true,
			KeyRight: true,
		},
	}
}

func (d *crossKeyDriver) handleKeyPressed(ctx context.Context, e KeyPressed) error {
	if ok := d.crossKeys[e.Key]; !ok {
		return d.keyHandler.handleKeyPressed(ctx, e)
	}

	if d.pressedKey == KeyUnknown {
		d.pressedKey = e.Key

		return d.keyHandler.handleKeyPressed(ctx, e)
	}

	d.nextPressedKey = e.Key

	return nil
}

func (d *crossKeyDriver) handleKeyReleased(ctx context.Context, e KeyReleased) error {
	if ok := d.crossKeys[e.Key]; !ok {
		return d.keyHandler.handleKeyReleased(ctx, e)
	}

	if d.pressedKey == e.Key {
		d.pressedKey = d.nextPressedKey
		d.nextPressedKey = KeyUnknown

		if err := d.keyHandler.handleKeyReleased(ctx, e); err != nil {
			return err
		}

		return d.keyHandler.handleKeyPressed(ctx, KeyPressed{
			Key: d.pressedKey,
		})
	}

	return nil
}

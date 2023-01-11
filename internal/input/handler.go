package input

import (
	"bombom/internal/pkg/event"

	"context"
)

type keyHandler interface {
	handleKeyPressed(context.Context, KeyPressed) error
	handleKeyReleased(context.Context, KeyReleased) error
}

func RegisterKeyHandlers(eventDispatcher *event.Dispatcher) {
	var aKeyHandler keyHandler

	aKeyHandler = &keyPrinter{}
	aKeyHandler = newCrossKeyFilter(aKeyHandler)

	eventDispatcher.RegisterHandler(
		event.New(KeyPressed{}).Type(),
		event.Handle[KeyPressed](aKeyHandler.handleKeyPressed),
	)

	eventDispatcher.RegisterHandler(
		event.New(KeyReleased{}).Type(),
		event.Handle[KeyReleased](aKeyHandler.handleKeyReleased),
	)
}

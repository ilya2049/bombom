package game

import (
	"bombom/internal/pkg/event"

	"context"
	"fmt"
)

func RegisterKeyHandlers(eventDispatcher *event.Dispatcher) {
	anInputEventsHandler := inputEventsHandler{}

	eventDispatcher.RegisterHandler(
		event.New(KeyPressed{}).Type(),
		event.Handle[KeyPressed](anInputEventsHandler.handleKeyPressed),
	)

	eventDispatcher.RegisterHandler(
		event.New(KeyReleased{}).Type(),
		event.Handle[KeyReleased](anInputEventsHandler.handleKeyReleased),
	)
}

type Key string

const (
	KeyUnknown Key = ""
	KeyUp      Key = "up"
	KeyDown    Key = "down"
	KeyLeft    Key = "left"
	KeyRight   Key = "right"
)

type KeyPressed struct {
	Key Key
}

type KeyReleased struct {
	Key Key
}

type inputEventsHandler struct {
	pressedKey Key
}

func (h *inputEventsHandler) handleKeyPressed(_ context.Context, e KeyPressed) error {
	if h.pressedKey == KeyUnknown {
		h.pressedKey = e.Key

		fmt.Println("key pressed: ", e.Key)
	}

	return nil
}

func (h *inputEventsHandler) handleKeyReleased(_ context.Context, e KeyReleased) error {
	if h.pressedKey == e.Key {
		h.pressedKey = KeyUnknown

		fmt.Println("key released: ", e.Key)
	}

	return nil
}

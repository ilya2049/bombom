package event

import (
	"context"
	"errors"
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

var ErrClarification = errors.New("failed to clarify an event")

type Handler interface {
	Handle(context.Context, Event) error
}

type Handle[T any] func(context.Context, T) error

func (h Handle[T]) Handle(ctx context.Context, e Event) error {
	encodedEvent, ok := e.(encoded)
	if ok {
		var (
			json           = jsoniter.ConfigFastest
			clarifiedEvent T
		)

		if err := json.Unmarshal(encodedEvent.data, &clarifiedEvent); err != nil {
			return fmt.Errorf("failed to unmarshal encoded event: %w", err)
		}

		return h(ctx, clarifiedEvent)
	}

	var rawEvent any
	rawEvent = e

	baseEvent, ok := e.(baseEvent)
	if ok {
		rawEvent = baseEvent.rawEvent
	}

	clarifiedEvent, ok := rawEvent.(T)
	if !ok {
		return fmt.Errorf("%w: %s", ErrClarification, e.Type())
	}

	return h(ctx, clarifiedEvent)
}

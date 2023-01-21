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
	switch typedEvent := e.(type) {
	case jsonSerializedEvent:
		return h.deserializeJSON(ctx, typedEvent)
	case autoTypedEvent:
		return h.deserializeBaseEvent(ctx, typedEvent)
	default:
		return h.deserialize(ctx, e)
	}
}

func (h Handle[T]) deserializeJSON(ctx context.Context, e jsonSerializedEvent) error {
	var (
		json           = jsoniter.ConfigFastest
		clarifiedEvent T
	)

	if err := json.Unmarshal(e.data, &clarifiedEvent); err != nil {
		return fmt.Errorf("failed to unmarshal encoded event: %w", err)
	}

	return h(ctx, clarifiedEvent)
}

func (h Handle[T]) deserializeBaseEvent(ctx context.Context, e autoTypedEvent) error {
	clarifiedEvent, ok := e.rawEvent.(T)
	if !ok {
		return fmt.Errorf("%w: %s", ErrClarification, e.typeName)
	}

	return h(ctx, clarifiedEvent)
}

func (h Handle[T]) deserialize(ctx context.Context, e Event) error {
	clarifiedEvent, ok := e.(T)
	if !ok {
		return fmt.Errorf("%w: %s", ErrClarification, e.Type())
	}

	return h(ctx, clarifiedEvent)
}

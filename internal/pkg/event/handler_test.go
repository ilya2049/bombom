package event_test

import (
	"bombom/internal/pkg/event"

	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandle_BaseEvent(t *testing.T) {
	var logString string

	handler := event.Handle[TestEvent](func(_ context.Context, e TestEvent) error {
		logString = e.TestField + " detected"

		return nil
	})

	err := handler.Handle(context.Background(), event.New(TestEvent{
		TestField: "test_field",
	}))
	require.NoError(t, err)

	assert.Equal(t, "test_field detected", logString)
}

func TestHandle_BaseEventWithEventPointer(t *testing.T) {
	var logString string

	handler := event.Handle[*TestEvent](func(_ context.Context, e *TestEvent) error {
		logString = e.TestField + " detected"

		return nil
	})

	err := handler.Handle(context.Background(), event.New(&TestEvent{
		TestField: "test_field",
	}))
	require.NoError(t, err)

	assert.Equal(t, "test_field detected", logString)
}

func TestHandle_EventWithCustomType(t *testing.T) {
	var logString string

	handler := event.Handle[*TestEventWithCustomType](func(_ context.Context, e *TestEventWithCustomType) error {
		logString = e.TestField + " detected"

		return nil
	})

	err := handler.Handle(context.Background(), &TestEventWithCustomType{
		TestField: "test_field",
	})
	require.NoError(t, err)

	assert.Equal(t, "test_field detected", logString)
}

func TestHandle_EventClarificationError(t *testing.T) {
	handler := event.Handle[TestEventWithCustomType](func(_ context.Context, e TestEventWithCustomType) error {
		return nil
	})

	err := handler.Handle(context.Background(), &TestEventWithCustomType{
		TestField: "",
	})

	assert.EqualError(t, err, "failed to clarify an event: TestEventWithType")
}

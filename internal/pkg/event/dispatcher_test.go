package event_test

import (
	"bombom/internal/pkg/event"

	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDispatcher_Dispatch_NoHandler(t *testing.T) {
	dispatcher := event.NewDispatcher()

	err := dispatcher.Dispatch(context.Background(), event.New(TestEvent{}))

	assert.EqualError(t, err, "event is not registered: TestEvent")
}

func TestDispatcher_Dispatch_OneHandler(t *testing.T) {
	dispatcher := event.NewDispatcher()

	var logString string

	handlerFunc := func(_ context.Context, e TestEvent) error {
		logString = e.TestField + " detected"

		return nil
	}

	dispatcher.RegisterHandler(event.New(TestEvent{}).Type(), event.Handle[TestEvent](handlerFunc))

	err := dispatcher.Dispatch(context.Background(), event.New(TestEvent{
		TestField: "test_field",
	}))
	require.NoError(t, err)

	assert.Equal(t, "test_field detected", logString)
}

func TestDispatcher_Dispatch_TwoHandlers(t *testing.T) {
	dispatcher := event.NewDispatcher()

	var logString string

	handlerOneFunc := func(_ context.Context, e TestEvent) error {
		logString += e.TestField + " detected by handler 1; "

		return nil
	}

	handlerTwoFunc := func(_ context.Context, e TestEvent) error {
		logString += e.TestField + " detected by handler 2"

		return nil
	}

	dispatcher.RegisterHandler(event.New(TestEvent{}).Type(), event.Handle[TestEvent](handlerOneFunc))
	dispatcher.RegisterHandler(event.New(TestEvent{}).Type(), event.Handle[TestEvent](handlerTwoFunc))

	err := dispatcher.Dispatch(context.Background(), event.New(TestEvent{
		TestField: "test_field",
	}))
	require.NoError(t, err)

	assert.Equal(t, "test_field detected by handler 1; test_field detected by handler 2", logString)
}

func TestDispatcher_Dispatch_ErrorOccurred(t *testing.T) {
	dispatcher := event.NewDispatcher()

	var logString string

	handlerOneFunc := func(_ context.Context, e TestEvent) error {
		logString += e.TestField + " detected only by handler 1"

		return nil
	}

	handlerTwoFunc := func(_ context.Context, e TestEvent) error {
		return ErrTest
	}

	dispatcher.RegisterHandler(event.New(TestEvent{}).Type(), event.Handle[TestEvent](handlerOneFunc))
	dispatcher.RegisterHandler(event.New(TestEvent{}).Type(), event.Handle[TestEvent](handlerTwoFunc))

	err := dispatcher.Dispatch(context.Background(), event.New(TestEvent{
		TestField: "test_field",
	}))
	require.EqualError(t, err, "failed to handle TestEvent by handler 2: test event error")

	assert.Equal(t, "test_field detected only by handler 1", logString)
}

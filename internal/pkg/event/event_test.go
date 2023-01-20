package event_test

import (
	"bombom/internal/pkg/event"

	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ErrTest = errors.New("test event error")

type TestEvent struct {
	TestField string `json:"test_field"`
}

type TestEventWithCustomType struct {
	TestField string
}

func (e *TestEventWithCustomType) Type() string {
	return "TestEventWithType"
}

func Test_baseEvent_Type(t *testing.T) {
	e := event.New(TestEvent{
		TestField: "test_field",
	})

	assert.Equal(t, "TestEvent", e.Type())
}

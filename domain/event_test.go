package domain_test

import (
	"context"
	"testing"
	"time"

	"github.com/quadroops/pkg/domain"
	"github.com/stretchr/testify/assert"
)

type FakeProvider struct {
	ch chan *domain.EventMessage
}

func NewFakeProvider() *FakeProvider {
	ch := make(chan *domain.EventMessage, 1)
	return &FakeProvider{ch}
}

func (fp *FakeProvider) Provide() <-chan *domain.EventMessage {
	return fp.ch
}

func (fp *FakeProvider) DoSomething(eventName, message string) {
	go func() {
		fp.ch <- &domain.EventMessage{
			Name:    domain.EventName(eventName),
			Payload: message,
		}
	}()
}

func TestEventHandlerSuccess(t *testing.T) {
	message := "hello world"
	provider := NewFakeProvider()

	handler := func(payload interface{}) {
		msg, valid := payload.(string)
		assert.True(t, valid)
		assert.Equal(t, "hello world", msg)
	}

	event := domain.NewEventProvider(provider)
	event.Setup().
		RegisterHandler("event.test", handler).
		Listen(context.TODO())

	provider.DoSomething("event.test", message)

	time.Sleep(1 * time.Second)
}

func TestEventHandlerMultiHandlers(t *testing.T) {
	message := "hello world"
	message2 := "hello world2"
	provider := NewFakeProvider()

	handler := func(payload interface{}) {
		msg, valid := payload.(string)
		assert.True(t, valid)
		assert.Equal(t, "hello world", msg)
	}

	handler2 := func(payload interface{}) {
		msg, valid := payload.(string)
		assert.True(t, valid)
		assert.Equal(t, "hello world", msg)
	}

	handler3 := func(payload interface{}) {
		msg, valid := payload.(string)
		assert.True(t, valid)
		assert.Equal(t, "hello world2", msg)
	}

	event := domain.NewEventProvider(provider)
	event.Setup().
		RegisterHandler("event.test", handler).
		RegisterHandler("event.test", handler2).
		RegisterHandler("event.test2", handler3).
		Listen(context.TODO())

	provider.DoSomething("event.test", message)
	provider.DoSomething("event.test2", message2)

	// this event will be ignored, because there are
	// no registered handlers
	provider.DoSomething("event.test3", message2)

	time.Sleep(1 * time.Second)
}

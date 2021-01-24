package domain

import (
	"context"

	"github.com/reactivex/rxgo/v2"
)

// EventName is a [Value Object] represent event's name
type EventName string

// EventMessage is main message object will passed to event's channel
type EventMessage struct {
	Name    EventName
	Payload interface{}
}

// EventHandler is a signature to proceed event's payload message
type EventHandler func(payload interface{})

// EventProvider should be implemented by all objects that provide
// event mechanism
type EventProvider interface {
	Provide() <-chan *EventMessage
}

// EventRunner is a main struct
type EventRunner struct {
	source    EventProvider
	eventChan <-chan *EventMessage
	handlers  map[EventName][]EventHandler
}

// NewEventProvider used to create new instance of EventProvider
// on this method, we just need to setup source
func NewEventProvider(provider EventProvider) *EventRunner {
	return &EventRunner{
		source:   provider,
		handlers: make(map[EventName][]EventHandler),
	}
}

// RegisterHandler used to register a handler and assign it with some event
func (ev *EventRunner) RegisterHandler(eventName EventName, handler EventHandler) *EventRunner {
	var mappers map[EventName][]EventHandler
	var handlers []EventHandler

	if registeredHandlers, exist := ev.handlers[eventName]; !exist {
		mappers = make(map[EventName][]EventHandler)
	} else {
		mappers = ev.handlers
		handlers = registeredHandlers
	}

	mappers[eventName] = append(handlers, handler)
	ev.handlers = mappers
	return ev
}

// Setup will call source's provider to fetch channel of event message
func (ev *EventRunner) Setup() *EventRunner {
	ev.eventChan = ev.source.Provide()
	return ev
}

// Listen start to proceed incoming events, this library only do an
// async and fire & forget mechanism, we doesn't need to wait to any
// response value
func (ev *EventRunner) Listen(ctx context.Context) {
	go func() {
		// listen to current event
		for event := range ev.eventChan {

			// only continue process for any events with registered handlers
			// an event which doesn't have any handlers will be ignored
			if handlers, exist := ev.handlers[event.Name]; exist {
				ch := make(chan rxgo.Item)
				go func() {
					ch <- rxgo.Of(event)
					close(ch)
				}()

				observable := rxgo.FromChannel(ch, rxgo.WithPublishStrategy())
				for _, handler := range handlers {
					observable.DoOnNext(func(i interface{}) {
						message, valid := i.(*EventMessage)
						if valid {
							// all handlers will be run async on different
							// goroutines
							go handler(message.Payload)
						}
					})
				}

				ctx, cancel := observable.Connect(ctx)

				<-ctx.Done()
				cancel()
			}

		}
	}()
}

# Domain

This is a simple library to supporting working using DDD (domain driven design).  This library provides:

- base entity
- simple event management

## Base Entity

```go
type Base struct {
	UID       UID       `json:"uid"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
```

You can embed this struct to another struct as base data, example:

```go
type MyEntity struct {
    *domain.Base
    Field1 string
    Field2 int
}

// usages
e := MyEntity{Field1: "testing", Field2: 10}

fmt.Println(e.UID)
fmt.Println(e.CreatedAt)
fmt.Println(e.UpdatedAt)
```

## Event Management

When we are working with DDD, we will need to manage a "domain event".  Sometimes after, we run a domain logic, we need to running an activity, which will produce some "side effect".

Actually, we can handle this just like procedural flows.  This library try to helping with provide more "clean" mechanism to manage domain event, by following reactive pattern, and actually this library internally using [RxGo](https://github.com/ReactiveX/RxGo) too to manage event's flows.

Each of entity which need to publish an event they have to implement `EventProvider` interface.

```go
// EventProvider should be implemented by all objects that provide
// event mechanism
type EventProvider interface {
	Provide() <-chan *EventMessage
}
```

It's just a simple interface to generate a channel of `EventMessage` .

```go
// EventMessage is main message object will passed to event's channel
type EventMessage struct {
	Name    EventName
	Payload interface{}
}
```

Example of an entity which has an event :

```go
// Login is main usecase object to manage login activity
type Login struct {
	repo             token.Repo
	tokenizer        token.Tokenizer
	repoCred         credential.Repo
	credentialHasher credential.Hasher
	uidGenerator     domain.UIDGenerator
	channelEvent     chan *domain.EventMessage
}

// NewLogin used create new instance of Login's struct
func NewLogin(
	repo token.Repo,
	repoCred credential.Repo,
	tokenizer token.Tokenizer,
	hasher credential.Hasher,
	uidGen domain.UIDGenerator) *Login {

	ch := make(chan *domain.EventMessage)
	return &Login{repo, tokenizer, repoCred, hasher, uidGen, ch}
}

// Provide used to giving channel of domain's event
func (l *Login) Provide() <-chan *domain.EventMessage {
	return l.channelEvent
}

func (l *Login) Auth(ctx context.Context) {
    // notify handlers
	go func(t *token.Token, cred *credential.Credential) {
		payload := token.EventPayload{
			CredentialUID: cred.UID,
			PublicKey:     cred.KeyPublic,
			AccessToken:   t.AccessToken,
			RefreshToken:  t.RefreshToken,
		}

		l.channelEvent <- &domain.EventMessage{
			Name:    token.TokenCreated,
			Payload: payload,
		}

		close(l.channelEvent)
	}(t, cred)
}
```

You can publish any event from anywhere in your code, as long as, you put the `EventMessage` into event's channel.

Just make sure, that you have to register your event's handler before you run your main logic.

A handler is just a simple function's signature like this:

```go
// EventHandler is a signature to proceed event's payload message
type EventHandler func(payload interface{})
```

Usages :

```go
login := usecase.NewLogin(
		repoToken,
		repoCred,
		buildTokenizer("token", nil),
		hasher("hashed", nil),
		uidGenerator(domain.UID("uid"), nil),
	)

event := domain.NewEventProvider(login)
event.Setup().RegisterHandler(token.TokenCreated, func(payload interface{}) {
    input, valid := payload.(token.EventPayload)
    assert.True(t, valid)
    assert.Equal(t, input.CredentialUID.String(), "test-token-refresh")
}).Listen(context.TODO())

// it will notify an event which should be catched by registered handler
resp, err := login.Auth(context.TODO(), "key", "secret", &usecase.AuthOption{})
```

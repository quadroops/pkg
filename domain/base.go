package domain

import (
	"time"
)

// UID is a [Value Object] represent some entity's unique identifier
type UID string

func (u UID) String() string {
	return string(u)
}

// Generator is main interface used to generate some value
type Generator interface {
	UID() (UID, error)
	RandomString() (string, error)
}

// EntityValidate is an interface which should be implemented by all entities
// that need to validated
type EntityValidate interface {
	Validate() error
}

// ValidateEntity used to validate any entities which implement EntityValidate
func ValidateEntity(e EntityValidate) error {
	return e.Validate()
}

// Base should be implemented for any entities that will saved on database
type Base struct {
	UID       UID       `json:"uid"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

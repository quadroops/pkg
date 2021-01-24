package domain_test

import (
	"errors"
	"testing"

	"github.com/quadroops/pkg/domain"
	"github.com/stretchr/testify/assert"
)

type FakeEntity struct {
	CondErr error
}

func (fe *FakeEntity) Validate() error {
	return fe.CondErr
}

func TestUIDToString(t *testing.T) {
	uid := domain.UID("uid")
	assert.Equal(t, "uid", uid.String())
}

func TestValidateEntityReturnNil(t *testing.T) {
	fe := FakeEntity{CondErr: nil}
	err := domain.ValidateEntity(&fe)
	assert.NoError(t, err)
}

func TestValidateEntityReturnError(t *testing.T) {
	fe := FakeEntity{CondErr: errors.New("testing")}
	err := domain.ValidateEntity(&fe)
	assert.Error(t, err)
	assert.Equal(t, "testing", err.Error())
}

package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/valueobject"
	"github.com/stretchr/testify/assert"
)
func TestNewPersonStatuses(t *testing.T) {
	SetupFiberTests(t)
	valueobject.InitPersonStatuses(os.Getenv("PERSON_STATUSES"))
	expected := int(2)
	actual := len(valueobject.GetPersonStatuses())
	message :=fmt.Sprintf("expected %d person statuses, got %d", expected, actual)
	assert.Equal(t, expected, len(valueobject.GetPersonStatuses()), message)
}
func TestPersonStatusesIsIn(t *testing.T) {
	SetupFiberTests(t)
	valueobject.InitPersonRoles(os.Getenv("PERSON_STATUSES"))
	assert.Equal(t, true, valueobject.IsInPersonRoles("Active"), "expected Active to be in PersonStatuses")
}
func TestPersonStatusesIsInFalse(t *testing.T) {
	SetupFiberTests(t)
	valueobject.InitPersonRoles(os.Getenv("PERSON_STATUSES"))
	assert.Equal(t, false, valueobject.IsInPersonRoles("ACTIVE"), "expected Active to be in PersonStatuses")
}
package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/valueobject"
	"github.com/stretchr/testify/assert"
)
func TestNewPersonRoles(t *testing.T) {
	SetupFiberTests(t)
	valueobject.InitPersonRoles(os.Getenv("PERSON_ROLES"))
	expected := int(5)
	actual := len(valueobject.GetPersonRoles())
	message :=fmt.Sprintf("expected %d person roles, got %d", expected, actual)
	assert.Equal(t, expected, len(valueobject.GetPersonRoles()), message)
}
func TestPersonRolesIsIn(t *testing.T) {
	SetupFiberTests(t)
	valueobject.InitPersonRoles(os.Getenv("PERSON_ROLES"))
	assert.Equal(t, true, valueobject.IsInPersonRoles("SYSTEM"), "expected SYSTEM to be in PersonRoles")
}
func TestPersonRolesIsInFalse(t *testing.T) {
	SetupFiberTests(t)
	valueobject.InitPersonRoles(os.Getenv("PERSON_ROLES"))
	assert.Equal(t, false, valueobject.IsInPersonRoles("SYSTEm"), "expected SYSTEM to be in PersonRoles")
}
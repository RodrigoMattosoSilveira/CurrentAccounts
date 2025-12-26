package valueobject

import "slices"

import "strings"

var personRoles = []string{}

// Initilized at bootstrap time
func InitPersonRoles(roles string) {
	personRoles = strings.Split(roles, ",")
}

func GetPersonRoles() []string {
	return personRoles
}

func IsInPersonRoles(candidate string) bool {
	return slices.Contains(personRoles, candidate)
}
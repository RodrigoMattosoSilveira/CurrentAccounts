package valueobject

import "strings"

var personRoles = []string{}

// Initilized at boorstrap time
func InitPersonRoles(roles string) {
	personRoles = strings.Split(roles, ",")
}

func GetPersonRoles() []string {
	return personRoles
}

func IsInPersonRoles(candidate string) bool {
	for _, role := range personRoles {
		if candidate == role {
			return true
		}
	}
	return false
}
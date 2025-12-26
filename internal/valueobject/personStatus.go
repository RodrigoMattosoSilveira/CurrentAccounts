package valueobject

import "slices"

import "strings"

var personStatuses []string

// Initilized at bootstrap time
func InitPersonStatuses(statuses string) {
	personStatuses = strings.Split(statuses, ",")
}

func GetPersonStatuses() []string {
	return personStatuses
}

func IsInPersonStatuses(candidate string) bool {
	return slices.Contains(personStatuses, candidate)
}
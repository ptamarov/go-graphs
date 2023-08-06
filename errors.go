package graph

import (
	"strings"
)

type manyErrors struct {
	errors []error
}

func (me manyErrors) Error() string {
	messages := []string{}

	for _, err := range me.errors {
		messages = append(messages, err.Error())
	}

	return strings.Join(messages, "\n")
}

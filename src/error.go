package turing

import "fmt"

type CypherNotFoundError string

func (e CypherNotFoundError) Error() string {
	return string(e)
}

type ResourceNotFoundError string

func (e ResourceNotFoundError) Error() string {
	return fmt.Sprintf("Resource '%v' not found", e)
}

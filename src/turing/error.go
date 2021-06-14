package turing

type CypherNotFoundError string

func (e CypherNotFoundError) Error() string {
	return string(e)
}

package turing

import (
	"fmt"

	"github.com/long-schlong-gang/turing/cypher"
)

var cypherRegistry = map[string]cypher.Cypher{
	"blank":  &cypher.BlankCypher{},
	"caesar": &cypher.CaesarCypher{},
}

func RegistryAddCypher(name string, c cypher.Cypher) {
	cypherRegistry[name] = c
}

func RegistryGetCypher(name string) (cypher.Cypher, error) {
	c, exists := cypherRegistry[name]
	if !exists {
		return nil, CypherNotFoundError(fmt.Sprintf("Cypher '%v' could not be found.", name))
	}

	return c, nil
}

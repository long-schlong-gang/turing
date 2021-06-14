package turing

import (
	"fmt"

	"github.com/long-schlong-gang/turing/cypher"
)

var cypherRegistry = map[string]Cypher{
	"blank":  cypher.BlankCypher{},
	"caesar": cypher.CaesarCypher{},
}

func RegistryAddCypher(name string, c Cypher) {
	cypherRegistry[name] = c
}

func RegistryGetCypher(name string) (Cypher, error) {
	c, exists := cypherRegistry[name]
	if !exists {
		return nil, CypherNotFoundError(fmt.Sprintf("Cypher '%v' could not be found.", name))
	}

	return c, nil
}

package cypher

/**
*	CAESAR CIPHER
*
*	Reference implementation to
*	test the framework
*
 */

import (
	"github.com/long-schlong-gang/turing"
)

type CaesarCypher struct {
	turing.Cypher
}

func (c *CaesarCypher) Encypher(plain string, key turing.Key) string {
	return plain
}

func (c *CaesarCypher) Decypher(plain string, key turing.Key) string {
	return plain + " - Romani ite domum!"
}

func (c *CaesarCypher) Crack(plain string) string {
	return plain
}

func (c *CaesarCypher) Confidence(plain string) float32 {
	return 0.0
}

func (c *CaesarCypher) KeyType() turing.KeyType {
	return turing.IntKey
}

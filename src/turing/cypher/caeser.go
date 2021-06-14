package cypher

/**
*	CAESAR CIPHER
*
*	Reference implementation to
*	test the framework
*
 */

// TODO: Actually implement this basic-ass cypher, lol

type CaesarCypher struct {
	Cypher
}

func (c *CaesarCypher) Encypher(plain string, key Key) string {
	return plain
}

func (c *CaesarCypher) Decypher(plain string, key Key) string {
	return plain + " - Romani ite domum!"
}

func (c *CaesarCypher) Crack(plain string) string {
	return plain
}

func (c *CaesarCypher) Confidence(plain string) float32 {
	return 0.0
}

func (c *CaesarCypher) KeyType() KeyType {
	return IntKey
}

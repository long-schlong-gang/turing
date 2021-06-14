package cypher

/**
*	BLANK CIPHER
*
*	Reference implementation to
*	show how the cyphers should
*	be layed out.
*
*	Just returns the plaintext as-is.
*
 */

type BlankCypher struct {
	Cypher
}

func (c *BlankCypher) Encypher(plain string, key Key) string {
	return plain
}

func (c *BlankCypher) Decypher(plain string, key Key) string {
	return plain
}

func (c *BlankCypher) Crack(plain string) string {
	return plain
}

func (c *BlankCypher) Confidence(plain string) float32 {
	return 0.0
}

func (c *BlankCypher) KeyType() KeyType {
	return NoKey
}

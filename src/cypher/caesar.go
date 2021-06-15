package cypher

/**
*	CAESAR CIPHER
*
*	Reference implementation to
*	test the framework
*
**/

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

/**
* Main Rotation Code adapted from
* github.com/alexjohnj/caesar
**/
func rotText(txt string, key int) string {
	key %= 26

	rotTxt := []byte(txt)

	alphabetLower := "abcdefghijklmnopqrstuvwxyz"
	alphabetUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for i, b := range rotTxt {
		if b >= 'a' && b <= 'z' {
			newI := (int(26+(b-'a')) + key) % 26
			rotTxt[i] = alphabetLower[newI]
		} else if b >= 'A' && b <= 'Z' {
			newI := (int(26+(b-'A')) + key) % 26
			rotTxt[i] = alphabetUpper[newI]
		}
	}
	return string(rotTxt)
}

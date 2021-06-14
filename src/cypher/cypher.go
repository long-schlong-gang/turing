package cypher

import "strconv"

type KeyType int

/// Types of Keys ///
const NoKey = KeyType(0)
const IntKey = KeyType(1)
const IntArrKey = KeyType(2)
const StringKey = KeyType(3)
const StringArrKey = KeyType(4)

type Key struct {
	Data []byte
	Type KeyType
}

var NilKey = Key{
	Data: []byte{},
	Type: NoKey,
}

func ParseKey(s string) Key {

	// Int Key Parse
	i, err := strconv.Atoi(s)
	if err == nil {
		return Key{
			Data: []byte{byte(i)},
			Type: IntKey,
		}
	}

	// TODO: Int Array key parse
	// TODO: String Array Parse

	// String Key Parse
	return Key{
		Data: []byte(s),
		Type: StringKey,
	}
}

type Cypher interface {
	Encypher(string, Key) string
	Decypher(string, Key) string
	Crack(string) string
	Confidence(string) float32
	KeyType() KeyType
}

// Idea: Make 'Cypher chain'/'pipeline' object that can be parsed from input to link multiple cyphers together (don't forget about keys tho)

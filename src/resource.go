package turing

/// Commonly used Resources ///
/* TODO: Should be replaced by a Database*/

var resource map[string]string = map[string]string{
	"ALPHABET_EN_UPPER": "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	"ALPHABET_EN_LOWER": "abcdefghijklmnopqrstuvwxyz",
}

// TODO: make resources interface{} and have funcs for retrieving specific types

func GetResource(key string) (string, error) {
	res, exists := resource[key]
	if !exists {
		return "", ResourceNotFoundError(res)
	}

	return res, nil
}

package main

// Dictionary is a data structure that associates a word with its definition
type Dictionary map[string]string

type wordNotFoundErr string

func (e wordNotFoundErr) Error() string {
	return string(e)
}

// ErrWordNotFound is used to signal the situation in which you try to find a
// word that is not available in the dictionary
const ErrWordNotFound = wordNotFoundErr("could not find the word you were looking for")

// Search returns the value associated to the given word in the dictionary
func (d Dictionary) Search(w string) (string, error) {
	def, ok := d[w]
	if !ok {
		return "", ErrWordNotFound
	}

	return def, nil
}
package main

// Dictionary is a data structure that associates a word with its definition
type Dictionary map[string]string

type dictionaryErr string

func (e dictionaryErr) Error() string {
	return string(e)
}

const (
	// ErrWordNotFound is used to signal the situation in which you try to find a
	// word that is not available in the dictionary
	ErrWordNotFound = dictionaryErr("could not find the word you were looking for")

	// ErrWordAlreadyExists is used to signal the situation in which you try to add
	// a word that is already present in the dictionary
	ErrWordAlreadyExists = dictionaryErr("word already exists")

	// ErrWordDoesNotExist is used to signal the situation in which you try to update
	// a word that is not present in the dictionary
	ErrWordDoesNotExist = dictionaryErr("word does not exist")
)

// Search returns the value associated to the given word in the dictionary
func (d Dictionary) Search(w string) (string, error) {
	def, ok := d[w]
	if !ok {
		return "", ErrWordNotFound
	}

	return def, nil
}

// Add includes a new word into the dictionary and associates it to the given definition
func (d Dictionary) Add(word string, def string) error {
	_, err := d.Search(word)
	switch err {
	case ErrWordNotFound:
		d[word] = def
	case nil:
		return ErrWordAlreadyExists
	default:
		return err
	}
	return nil
}

// Update amends the definition of an existing word
func (d Dictionary) Update(word string, newDef string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = newDef
	case ErrWordNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}
	return nil
}

// Delete removes a word from the dictionary
func (d Dictionary) Delete(w string) {
	delete(d, w)
}



package gordle

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type corpusError string

func (err corpusError) Error() string {
	return string(err)
}

// ErrCorpusIsEmpty is a sentinel error signaling the corpus file is empty
const ErrCorpusIsEmpty = corpusError("corpus file is empty")

// ReadCorpus reads the file identified by the given path, which should be a
// newline separated list of words, which is returned by this function.
func ReadCorpus(path string) ([]string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Unable to open %q for reading: %w", path, err)
	}

	if len(bytes) == 0 {
		return nil, ErrCorpusIsEmpty
	}

	// splits a string by the blanks
	words := strings.Fields(string(bytes))
	return words, nil
}

// pickWord takes a corpus and return a random word
func pickWord(corpus []string) string {
	rand.Seed(time.Now().UTC().UnixNano()) // this is no longer needed, but keeping it for reference
	index := rand.Intn(len(corpus))
	return corpus[index]
}

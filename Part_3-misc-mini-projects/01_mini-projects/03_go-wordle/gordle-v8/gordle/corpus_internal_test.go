package gordle

import "testing"

func TestPickWord(t *testing.T) {
	corpus := []string{"HELLO", "JASON", "IDRIS", "JONES"}
	for i := 0; i < 100; i++ {
		randWord := pickWord(corpus)
		if !inCorpus(corpus, randWord) {
			t.Errorf("iteration #%v: expected %q to be in corpus %v, but wasn't", i, randWord, corpus)
		}
	}
}

func inCorpus(corpus []string, word string) bool {
	for _, corpusWord := range corpus {
		if corpusWord == word {
			return true
		}
	}
	return false
}
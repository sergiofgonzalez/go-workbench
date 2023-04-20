package main

import "testing"

func TestSearch(t *testing.T) {
	t.Run("word found", func(t *testing.T) {
		dictionary := Dictionary{
			"test": "this is the value associated to test",
		}

		got, err := dictionary.Search("test")
		assertDefinition(t, got, "this is the value associated to test")
		assertError(t, err, nil)
	})

	t.Run("word not found", func(t *testing.T) {
		dictionary := Dictionary{
			"test": "this is the value associated to test",
		}

		got, err := dictionary.Search("nonexistent")
		assertDefinition(t, got, "")
		assertError(t, err, ErrWordNotFound)
	})
}


func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		errAdd := dictionary.Add("test", "this is the value associated to test")
		got, err := dictionary.Search("test")
		assertDefinition(t, got, "this is the value associated to test")
		assertError(t, err, nil)
		assertError(t, errAdd, nil)
	})

	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{
			"test": "this is the value associated to test",
		}
		err := dictionary.Add("test", "already existing")
		assertError(t, err, ErrWordAlreadyExists)
	})

}

func assertDefinition(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		if got != want {
			t.Errorf("unexpected definition: got %q, but wanted %q", got, want)
		}
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got != want {
		t.Errorf("unexpected error: got %v, but wanted %v", got, want)
	}
}
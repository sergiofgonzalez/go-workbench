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

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{
			"test": "this is the value associated to test",
		}
		errUpdate := dictionary.Update("test", "new definition")
		got, errSearch := dictionary.Search("test")
		assertError(t, errUpdate, nil)
		assertError(t, errSearch, nil)
		assertDefinition(t, got, "new definition")
	})

	t.Run("nonexisting word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Update("test", "new definition")
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{
			"test": "this is the value associated to test",
		}
		dictionary.Delete("test")
		_, err := dictionary.Search("test")
		assertError(t, err, ErrWordNotFound)
	})

	t.Run("nonexisting word", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Delete("test")
		_, err := dictionary.Search("test")
		assertError(t, err, ErrWordNotFound)
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